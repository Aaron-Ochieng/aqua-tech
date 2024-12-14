#include <WiFi.h>

HardwareSerial sim900(2); // Use UART2

// SIM900 module pins
#define RX_PIN 16
#define TX_PIN 17
#define SIM900_BAUD 9600

// Replace with your APN, username, and password (if required)
const char apn[] = "your_apn";
const char user[] = ""; // Leave blank if not required
const char pass[] = ""; // Leave blank if not required

// Test API endpoint
const char* server = "jsonplaceholder.typicode.com";
const char* endpoint = "/posts";

// Sensor pin
#define SENSOR_PIN A0

void setup() {
  Serial.begin(115200);
  sim900.begin(SIM900_BAUD, SERIAL_8N1, RX_PIN, TX_PIN);

  Serial.println("Initializing SIM900...");
  sendCommand("AT", 1000);
  sendCommand("AT+CGATT=1", 1000);

  // Setup GPRS
  Serial.println("Setting up GPRS...");
  sendCommand("AT+SAPBR=3,1,\"CONTYPE\",\"GPRS\"", 1000);
  sendCommand("AT+SAPBR=3,1,\"APN\",\"" + String(apn) + "\"", 1000);
  sendCommand("AT+SAPBR=1,1", 5000);
  sendCommand("AT+SAPBR=2,1", 1000);
}

void loop() {
  float sensorValue = analogRead(SENSOR_PIN);
  float voltage = (sensorValue / 1023.0) * 5.0;

  String postData = "{\"sensor\": \"temperature\", \"value\": " + String(voltage) + "}";

  // Make POST request
  sendCommand("AT+HTTPINIT", 1000);
  sendCommand("AT+HTTPPARA=\"CID\",1", 1000);
  sendCommand("AT+HTTPPARA=\"URL\",\"http://" + String(server) + endpoint + "\"", 1000);
  sendCommand("AT+HTTPPARA=\"CONTENT\",\"application/json\"", 1000);
  sendCommand("AT+HTTPDATA=" + String(postData.length()) + ",10000", 1000);

  delay(100);
  sim900.println(postData);
  delay(1000);
  sendCommand("AT+HTTPACTION=1", 1000);
  delay(5000); // Wait for response

  sendCommand("AT+HTTPREAD", 1000);
  sendCommand("AT+HTTPTERM", 1000);

  delay(30000); // Wait before sending next data
}

void sendCommand(String command, int timeout) {
  sim900.println(command);
  long int time = millis();
  while ((time + timeout) > millis()) {
    while (sim900.available()) {
      char c = sim900.read();
      Serial.print(c);
    }
  }
}
