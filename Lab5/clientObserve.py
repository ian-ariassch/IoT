# client_observe.py

import logging
import asyncio

from aiocoap import *

logging.basicConfig(level=logging.INFO) 

def observe_callback(response): 
    if response.code.is_successful():
        print("Alarm status: %s" % (response.payload.decode('ascii')))
    else:
        print('Error code %s' % response.code)

async def main():
    context = await Context.create_client_context() #Se crea el contexto del cliente

    request = Message(code=GET) #Se crea un request de tipo mensaje con el método GET
    request.set_request_uri('coap://localhost/alarm') #Se le asigna el endpoint al request
    request.opt.observe = 0 
    observation_is_over = asyncio.Future() 

    try:
        context_request = context.request(request) #Se le asigna el request al contexto creado anteriormente
        context_request.observation.register_callback(observe_callback) #Se registra el callback para el request
                                                                        # este se llamará cada vez que se observe un cambio en el recurso
                                                                        # observado
        response = await context_request.response #Se espera a que se reciba una respuesta
        exit_reason = await observation_is_over
        print('Observation is over: %r' % exit_reason)
    finally:
        if not context_request.response.done(): #Si la respuesta no se ha recibido, se cancela el request
            context_request.response.cancel()
        if not context_request.observation.cancelled: #Si la observación no se ha cancelado, se cancela
            context_request.observation.cancel()

if __name__ == "__main__":
    asyncio.get_event_loop().run_until_complete(main())