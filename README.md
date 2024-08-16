# TalanaKombat

Esta es una prueba para Talana, esta hecho en GO + Arquitectura hexagonal. 

Crear un archivo .env en root.

'''
ENV= [dev|prod]
PORT= [Port for your app]
'''

Tienes una collecion con los request asociados a los enunciados de la prueba. Se corrigio el primero que presentaba una diferencia en los mensajes de salida con el input de ejemplo.


### Instalar dependencias

'''
Make install
'''

### Para correr local

'''
Make run
'''

### Para crear Docker

'''
Make docker-build
'''

### Para correr Docker

''
Make docker-run
'''