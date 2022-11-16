# client_put.py
import asyncio
import random
from aiocoap import *

async def main():
    context = await Context.create_client_context() #Se crea el contexto del cliente
    alarm_state = random.choice([True, False]) #Se elige un estado aleatorio para actualizar el recurso de la alarma
    print('Alarm state: %s' % alarm_state) 
    payload = b"OFF" 

    if alarm_state:
        payload = b"ON"

    print(payload)
    request = Message(code=PUT, payload=payload, uri="coap://localhost/alarm") #Se crea un request de tipo mensaje con el método PUT
                                                                            # en el endpoint del recurso, en este caso alarm.

    response = await context.request(request).response #Si es exitoso el request, se actualiza el recurso de la alarma en el servidor 
                                                       # y se obtiene código 2.04. 
    print('Result: %s\n%r'%(response.code, response.payload))
    
if __name__ == "__main__":
    asyncio.get_event_loop().run_until_complete(main())