import org.apache.spark.sql.SparkSession
import org.apache.spark.sql.functions._
import org.apache.spark.sql.streaming.Trigger
import org.apache.spark.sql.types.{StructType, StringType, TimestampType}

object SparkStructuredStreamingApp {
  def main(args: Array[String]): Unit = {
    val spark = SparkSession.builder()
      .appName("ClickEventStreaming")
      .master("local[*]") // Change to your Spark cluster URL for deployment
      .getOrCreate()

    import spark.implicits._

    val schema = new StructType()
      .add("sessionId", StringType)
      .add("eventType", StringType)
      .add("ts", TimestampType)

    val clickStreamDF = spark.readStream
      .format("kafka")
      .option("kafka.bootstrap.servers", "localhost:9092") // Kafka server
      .option("subscribe", "<your_name>/clickstream")
      .load()
      .selectExpr("CAST(value AS STRING) as json")
      .select(from_json($"json", schema).as("data"))
      .select($"data.*")

    val kpi1DF = clickStreamDF
      .filter($"eventType" === "FormSubmit")
      .groupBy(window($"ts", "1 hour"), $"sessionId")
      .agg(
        avg($"ts").alias("avgSubmitTime"),
        min($"ts").alias("minSubmitTime"),
        max($"ts").alias("maxSubmitTime")
      )

    val kpi2DF = clickStreamDF
      .filter($"eventType" === "Open")
      .groupBy(window($"ts", "15 minutes"))
      .agg(countDistinct($"sessionId").alias("activeSessions"))

    val kpi3DF = clickStreamDF
      .groupBy($"sessionId")
      .agg(
        collect_set($"eventType").alias("eventTypes")
      )
      .filter(
        !$"eventTypes".contains("FormSubmit") &&
        $"eventTypes".contains("Open") &&
        $"eventTypes".contains("Close")
      )
      .groupBy()
      .agg(count("*").alias("sessionsWithoutFormSubmit"))

    val query = kpi1DF.writeStream
      .outputMode("update")
      .trigger(Trigger.ProcessingTime("1 hour"))
      .format("delta")
      .option("checkpointLocation", "/path/to/checkpoint") // Change to a suitable checkpoint location
      .start("/path/to/delta/kpi1")

    val query2 = kpi2DF.writeStream
      .outputMode("update")
      .trigger(Trigger.ProcessingTime("15 minutes"))
      .format("delta")
      .option("checkpointLocation", "/path/to/checkpoint2") // Change to a suitable checkpoint location
      .start("/path/to/delta/kpi2")

    val query3 = kpi3DF.writeStream
      .outputMode("update")
      .trigger(Trigger.ProcessingTime("15 minutes"))
      .format("delta")
      .option("checkpointLocation", "/path/to/checkpoint3") // Change to a suitable checkpoint location
      .start("/path/to/delta/kpi3")

    query.awaitTermination()
    query2.awaitTermination()
    query3.awaitTermination()

    spark.stop()
  }
}
