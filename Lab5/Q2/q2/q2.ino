#include <ESP8266WiFi.h> 
#include <WiFiUdp.h>
#include <coap-simple.h>

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

const char *ssid = "XXXXXX"; // Enter your WiFi name
const char *password = "XXXXXXX"; // Enter WiFi password


WiFiClient espClient;
WiFiUDP Udp;

Coap coap(Udp);

void connectToWifi()
{
   WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
   delay(500);
   Serial.println("Connecting to WiFi..");
  }
  Serial.println("Connected to the WiFi network");
}

void callback_response(CoapPacket &packet, IPAddress ip, int port) //Definición de función que se ejecutará cuando el servidor responda.
{
  Serial.println("Server responded"); //No es necesario que realice una operación específica, ya que la librería se encarga de procesar la respuesta.
}
 
void IRAM_ATTR pulseCounter()
{
  pulseCount++;
}
 
 
void setup()
{
  Serial.begin(9600);

  connectToWifi(); //Funciones descritas en la pregunta 1
 
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

  coap.response(callback_response);

  coap.start();
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
    Serial.println(""); 

    String message = ""; //Construcción del mensaje y cast a char[] para que pueda ser enviado como payload en el publish.
    message += totalMilliLitres;

     char charBuf[50];
     message.toCharArray(charBuf,50);

    int messageId = coap.put(IPAddress(192,168,1,20), 5683, "water", charBuf);  //Envío del mensaje al servidor a través del método PUT.
                                                                                   //El primer parámetro es la IP del servidor.
                                                                                    //El segundo parámetro es el puerto en donde se encuentra el servidor.
                                                                                    //El tercer parámetro es el recurso al que se quiere actualizar.
                                                                                    //El cuarto parámetro es el mensaje a enviar, en este caso, el volumen de agua.
    Serial.print(messageId); //La operación anterior retorna un ID de mensaje, el cuál se imprime en el monitor serial.
    coap.loop();
 
  }
}
