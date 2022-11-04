#include <ESP8266WiFi.h>
#include <PubSubClient.h>// WiFi
const char *ssid = ""; // Enter your WiFi name
const char *password = ""; // Enter WiFi password
const char *mqtt_broker = ""; // IP del Broker
const char *topic = "test/topic"; //El topic al cual el arduino se suscribe. En el caso de este laboratorio podría ser el topic humedad.
const int mqtt_port = 1883;

WiFiClient espClient;
PubSubClient client(espClient);

void setup() {

  Serial.begin(115200);

  WiFi.begin(ssid, password);

while (WiFi.status() != WL_CONNECTED) {
 delay(500);
 Serial.println("Connecting to WiFi..");
}
Serial.println("Connected to the WiFi network");

client.setServer(mqtt_broker, mqtt_port); //Definir el servidor al broker utilizando su IP y el puerto.
//Con esta información el cliente (arduino) podrá saber a que dirección apuntar para la comunicación con el broker.

client.setCallback(callback); //Define la función a llamar cuando se recibe un mensaje en alguna de las suscripciones. 

 while (!client.connected()) {
    String client_id = "esp8266-client-"; //Se define el id del cliente. Este ID le servirá al broker para guardar sesiones en caso sean persistentes.
    client_id += String(WiFi.macAddress());
    Serial.printf("The client %s connects to mosquitto mqtt broker\n", client_id.c_str());
    if (client.connect(client_id.c_str())) { //Connecta al arduino al broker enviando un paquete MQTT "CONNECT", que contiene el id del cliente
      //e información de autenticación en caso sea necesaria.
      // Al el cliente enviar el paquete "CONNECT", el broker retorna un paquete MQTT "CONNACK" con información de si tiene la sesión del cliente guardada y 
      // un código de retorno indicando el estado de la conexión (exitosa, fallida y razón).
      Serial.println("Public emqx mqtt broker connected");
    } else {
    Serial.print("failed with state ");
    Serial.print(client.state()); //La razón de falla que es recibida en el paquete "CONNACK", se imprime aquí.
    delay(2000);
    }
  }
// publish and subscribe
 client.publish(topic, "Hello From ESP8266!"); //Envía información sobre al broker en un topic determinado. En este caso, esta envíando un mensaje al topic "test/topic".
 //Esto es realizado a través de un paquete MQTT llamado "Publish", este contiene la siguiente información:
 //PacketId: Identificador del paquete
 //Topic: Nombre del topic
 //QoS: Calidad de servicio, indica la garantía de que la información llegue a un suscriptor
 //Retain: Indica si el mensaje debe ser retenido para ser enviado a nuevos suscriptores.
 //Payload: El mensaje
 //DUP Flag: Un flag que indica si se ha intentado enviar el paquete más de una vez en caso el receptor no haya confirmado que lo recibió.
  
 client.subscribe(topic); //Suscribe al arduino al topic definido. Esto le hace saber al broker que el cliente (arduino) está interesado en recibir
 //información sobre este topic y le enviará datos en cuanto el broker reciba un publish de algún sensor correspondiente al topic.
 //Esto es realizado a través de un paquete MQTT "SUBSCRIBE" que contiene la siguiente información:
 //PacketId: Identificador del paquete
 //Topic: Nombre del topic a suscribirse
 //QoS: Calidad de servicio en la que deseen recibir la información.

 //Un paquete puede tener varias suscripciones, es decir, se puede enviar una lista de topics y QoS.

 //El broker al recibir un paquete SUBSCRIBE, retorna un paquete SUBACK con su PacketId y códigos de retorno para cada topic suscrito.
}

void callback(char *topic, byte *payload, unsigned int length) { //Esta función es ejecutada al recibir un mensaje de alguna suscripción.
  //Una función de callback para MQTT debe tener los parámetros topic, para saber a que topic le pertenece la publicación; el payload, el contenido del mensaje en sí
  //y el tamaño del payload para poder acceder a el sin dificultad.
  //La función realiza una impresión en la consola serial.
 Serial.print("Message arrived in topic: ");
  Serial.println(topic);
 Serial.print("Message:");
 for (int i = 0; i < length; i++) {
   Serial.print((char) payload[i]);
  }
  Serial.println();
  Serial.println(" - - - - - - - - - - - -");
}

void loop() {
  client.loop();
}
