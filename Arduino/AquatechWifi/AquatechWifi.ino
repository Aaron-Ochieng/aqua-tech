#include <WiFi.h>
#include <HTTPClient.h>
#include <DHT.h>

DHT dht(2, DHT11);

// WiFi credentials
const char* ssid = "Redmi Note 11S";
const char* password = "123456789";

// Server endpoint (modifiable through Serial Monitor)
String server = "https://grumpy-islands-rest.loca.lt/data";

void setup() {
  Serial.begin(115200);
  dht.begin();
  delay(2000);

  // Connect to WiFi
  Serial.println("Connecting to WiFi...");
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(1000);
    Serial.print(".");
  }
  Serial.println("\nWiFi connected");
  Serial.print("IP Address: ");
  Serial.println(WiFi.localIP());

  // Prompt for server configuration via Serial Monitor
  Serial.println("Current server: " + server);
  Serial.println("Enter a new server URL or press Enter to keep the current one:");
}

void loop() {
  // Check if Serial Monitor input is available
  if (Serial.available() > 0) {
    String input = Serial.readStringUntil('\n');
    input.trim(); // Remove extra whitespace and newline characters
    if (input.length() > 0) {
      server = input; // Update the server URL
      Serial.println("Server updated to: " + server);
    } else {
      Serial.println("Keeping the current server: " + server);
    }
  }

  // Read temperature and humidity from the DHT sensor
  float temp = dht.readTemperature();
  float humidity = dht.readHumidity();

  Serial.print("Temp: ");
  Serial.print(temp);
  Serial.println(" C");

  Serial.print("Humidity: ");
  Serial.print(humidity);
  Serial.println(" %");

  // Create JSON payload
  String postData = "{\"temp\": " + String(temp) + ", \"Humidity\": " + String(humidity) + "}";

  if (WiFi.status() == WL_CONNECTED) { // Ensure WiFi is connected
    HTTPClient http;

    Serial.println("Making POST request...");
    http.begin(server);
    http.addHeader("Content-Type", "application/json");

    int httpResponseCode = http.POST(postData);

    if (httpResponseCode > 0) {
      String response = http.getString();
      Serial.print("HTTP Response code: ");
      Serial.println(httpResponseCode);
      Serial.println("Response: ");
      Serial.println(response);
    } else {
      Serial.print("Error in HTTP request: ");
      Serial.println(httpResponseCode);
    }

    http.end(); // Free resources
  } else {
    Serial.println("WiFi not connected");
  }

//  delay(30000); // Wait before sending next data
}
