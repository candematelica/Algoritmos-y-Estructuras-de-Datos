#!/usr/bin/python3
from comandos import leer_comandos
from func_archivos import *

def main():
   archivoPaginasV = leer_archivos()
   g = cargar_vertices(archivoPaginasV)
   archivoPaginasV.close()
   
   archivoPaginasA = leer_archivos()
   cargar_aristas(g, archivoPaginasA)
   archivoPaginasA.close()
   
   leer_comandos(g)
    
if __name__ == "__main__":
   main()
