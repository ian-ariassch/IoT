import aiocoap.resource as resource
import aiocoap
import threading
import logging
import asyncio

class WaterResource(resource.ObservableResource):
    def __init__(self):
        super().__init__()

        self.totalWater = 0 
        self.has_observers = False 
        self.notify_observers = False 

    def notify_observers_check(self):   
        while True:
            if self.has_observers and self.notify_observers:
                print('Notifying observers')
                self.updated_state() 
                self.notify_observers = False 

    def update_observation_count(self, count): 
        if count:  
            self.has_observers = True
        else:
            self.has_observers = False

    async def render_get(self, request): 
        print('Return total water: %s' % self.totalWater) 
        payload = b'%s' % str(self.totalWater).encode('ascii')

        return aiocoap.Message(payload=payload) 

    async def render_put(self, request): 
        self.totalWater = request.payload.decode('ascii')
        print('Updated total water: %s' % self.totalWater)
        self.notify_observers = True

        return aiocoap.Message(code=aiocoap.CHANGED, payload=b'%s' % self.totalWater.encode('ascii')) 

logging.basicConfig(level=logging.INFO)
logging.getLogger("coap-server").setLevel(logging.DEBUG)

def main():
    root = resource.Site() 
    waterResource = WaterResource() 
    root.add_resource(['water'], waterResource) 
    asyncio.Task(aiocoap.Context.create_server_context(root, bind=('192.168.1.20', 5683))) 

    observers_notifier = threading.Thread(target=waterResource.notify_observers_check)
    observers_notifier.daemon = True
    observers_notifier.start()

    asyncio.get_event_loop().run_forever()

if __name__ == "__main__":
    main()