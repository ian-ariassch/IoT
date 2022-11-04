#include <ESP8266WiFi.h> 
#include <PubSubClient.h>

#define SENSOR D3

long currentMillis = 0;
long previousMillis = 0;
int interval = 1000;
float calibrationFactor = 4.5;
volatile byte pulseCount;
byte pulse1Sec = 0;
float flowRate;
unsigned long flowMilliLitres;
unsigned int totalMilliLitres;

const char *ssid = "SENNA"; // Enter your WiFi name
const char *password = "55555555"; // Enter WiFi password
const char *mqtt_broker = "192.168.1.20"; // IP del Broker
const char *topic = "water/militotal"; 
const int mqtt_port = 1883;

WiFiClient espClient;
PubSubClient client(espClient);

void connectToWifi()
{
   WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
   delay(500);
   Serial.println("Connecting to WiFi..");
  }
  Serial.println("Connected to the WiFi network");
}

void setupMQTTBroker()
{
  client.setServer(mqtt_broker, mqtt_port);

 while (!client.connected()) {
    String client_id = "esp8266-client-"; 
    client_id += String(WiFi.macAddress());
    Serial.printf("The client %s connects to mosquitto mqtt broker\n", client_id.c_str());
    if (client.connect(client_id.c_str())) { 
      Serial.println("Public emqx mqtt broker connected");
    } else {
    Serial.print("failed with state ");
    Serial.print(client.state()); 
    delay(2000);
    }
  }
}
 
void IRAM_ATTR pulseCounter()
{
  pulseCount++;
}
 
 
void setup()
{
  Serial.begin(9600);

  connectToWifi();
  setupMQTTBroker();
 
  pinMode(SENSOR, INPUT_PULLUP);
 
  pulseCount = 0;
  flowRate = 0.0;
  flowMilliLitres = 0;
  totalMilliLitres = 0;
  previousMillis = 0;
 
  attachInterrupt(digitalPinToInterrupt(SENSOR), pulseCounter, FALLING);
}
 
void loop()
{
  currentMillis = millis();
  if (currentMillis - previousMillis > interval) 
  {
    
    pulse1Sec = pulseCount;
    pulseCount = 0;
 
    flowRate = ((1000.0 / (millis() - previousMillis)) * pulse1Sec) / calibrationFactor;
    previousMillis = millis();
 
    flowMilliLitres = (flowRate / 60) * 1000;
 
    totalMilliLitres += flowMilliLitres;
     
    Serial.print("Output Liquid Quantity: ");
    Serial.print(totalMilliLitres);
    Serial.print("mL");

    String message = "Ouput Liquid Quantity: ";
    message += totalMilliLitres;

     char charBuf[50];
     message.toCharArray(charBuf,50);

    client.publish(topic, charBuf);
    
 
  }
}
