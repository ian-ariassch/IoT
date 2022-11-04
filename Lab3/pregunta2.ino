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
const char *password = "-"; // Enter WiFi password
const char *mqtt_broker = "192.168.1.20"; // IP del Broker
<<<<<<< HEAD
const char *topic = "water/militotal";
=======
const char *topic = "water/militotal"; 
>>>>>>> 356e3d374a433ff8b982ae357457e19c4be777a9
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

  connectToWifi(); //Funciones descritas en la pregunta 1
  setupMQTTBroker();
 
  pinMode(SENSOR, INPUT_PULLUP);
 
  pulseCount = 0;            //Declaración de variables iniciales
  flowRate = 0.0;
  flowMilliLitres = 0;
  totalMilliLitres = 0;
  previousMillis = 0;
 
  attachInterrupt(digitalPinToInterrupt(SENSOR), pulseCounter, FALLING); //Declaración de cuando ocurre un "Interrupt"
  //El primer parámetro es el PIN en donde se leerá la corriente. En este caso el input de sensor de flujo de agua.
  //El segundo parámetro es la función a llamar cuando ocurre un interrupt. En este caso, se busca sumar la cantidad de pulsos eléctricos cada vez que ocurre una interrupción.
  //El tercer parámetros define cuando ocurrirá una interrupción. En este caso, esta definido como "FALLING" que es cuando el pin va de "HIGH" a "LOW". 
}
 
void loop()
{
  currentMillis = millis();
  if (currentMillis - previousMillis > interval) //Se revisa si es que ha pasado más de 1 segundo.
  {
    
    pulse1Sec = pulseCount; //Se obtiene cuantos pulsos hubo en este segundo y se resetean.
    pulseCount = 0;
 
    flowRate = ((1000.0 / (millis() - previousMillis)) * pulse1Sec) / calibrationFactor; //Se calcula el flujo de agua. Se divide 1000 entre el tiempo que ha pasado desde la última ejecución en caso no se haya
    //completado en exactamente 1 segundo.
    //Luego, se multiplica ese valor por la cantidad de pulsos obtenidos en este tiempo. Finalmente, se divide entre el calibrationFactor, el cuál es un valor proporcionado por el fabricante del sensor.
    previousMillis = millis(); 
 
    flowMilliLitres = (flowRate / 60) * 1000; //Se divide entre 60 para obtener litros/segundo y luego se multiplica por 1000 para pasarlo a mililitros.
 
    totalMilliLitres += flowMilliLitres; //Por último, se suma el flujo al total de mililitros.
     
    Serial.print("Output Liquid Quantity: ");
    Serial.print(totalMilliLitres);
    Serial.print("mL");

    String message = "Ouput Liquid Quantity: "; //Construcción del mensaje y cast a char[] para que pueda ser enviado como payload en el publish.
    message += totalMilliLitres;

     char charBuf[50];
     message.toCharArray(charBuf,50);

    client.publish(topic, charBuf);
    
 
  }
}
