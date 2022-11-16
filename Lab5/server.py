import aiocoap.resource as resource
import aiocoap
import threading
import logging
import asyncio

class AlarmResource(resource.ObservableResource): #Creación de clase AlarmResource que hereda de ObservableResource. Es decir, 
    def __init__(self):
        super().__init__()

        self.status = "OFF" #Inicialización de estado de alarma
        self.has_observers = False #Inicialización de estado de observadores
        self.notify_observers = False #Inicialización de estado de notificación de observadores
                                      # Ambos inician apagados pues no hay observadores al iniciar el servidor

    def notify_observers_check(self):    #Servicio que se ejecuta en un hilo aparte para revisar si se debe notificar a los observadores
        while True:
            if self.has_observers and self.notify_observers:
                print('Notifying observers')
                self.updated_state()  #Función que notifica a los observadores de un cambio de estado
                self.notify_observers = False #Se reinicia el estado de notificación de observadores

    def update_observation_count(self, count): #Función que se ejecuta cuando se agrega o elimina un observador en el recurso
        if count:  #Si hay observadores se actualiza el flag, lo que causa que la función notify_observers_check notifique a los observadores
            self.has_observers = True
        else:
            self.has_observers = False

    async def render_get(self, request): #Implementación para consultas GET. Se retorna el payload con el estado de la alarma en ASCII
        print('Return alarm state: %s' % self.status) #Esta función también se ejecuta al notificar a los observadores con la función updated_state
        payload = b'%s' % self.status.encode('ascii')

        return aiocoap.Message(payload=payload) #Se envía un código de respuesta de 2.05 Content

    async def render_put(self, request): #Implementación para consultas PUT. Se actualiza el estado de la alarma y se enciende el flag de notificación a los observadores
        self.status = request.payload.decode('ascii')
        print('Update alarm state: %s' % self.status)
        self.notify_observers = True

        return aiocoap.Message(code=aiocoap.CHANGED, payload=b'%s' % self.status.encode('ascii')) #Se envía un código de respuesta de 2.04 Changed con el cambio de estado de la alarma en ASCII

logging.basicConfig(level=logging.INFO)
logging.getLogger("coap-server").setLevel(logging.DEBUG)

def main():
    root = resource.Site() #Creación del sitio raíz, es decir, el sitio que contiene todos los recursos. En este caso solo contiene
                            # el recurso AlarmResource que se creó anteriormente.
    alarmResource = AlarmResource() #Creación de objeto de tipo AlarmResource
    root.add_resource(['alarm'], alarmResource) #Se agrega el recurso AlarmResource al sitio raíz
    asyncio.Task(aiocoap.Context.create_server_context(root, bind=('localhost', 5683))) #Se crea el contexto del servidor y se le asigna el sitio raíz

    observers_notifier = threading.Thread(target=alarmResource.notify_observers_check) #Se crea un hilo para ejecutar la función notify_observers_check
    observers_notifier.daemon = True
    observers_notifier.start()

    asyncio.get_event_loop().run_forever()

if __name__ == "__main__":
    main()