import sys

from grafo import *

CANT_PARAMETROS = 2
ARCHIVO_PAGINAS = 1

def cargar_grafo(archivoPaginas):

   g = Grafo(es_dirigido=True)
   for pagina in archivoPaginas:
      p = pagina.strip().split("\n")
      if len(p) > 1:
         for link in p[1:]:
            if link not in g.obtener_vertices():
               g.agregar_vertice(link)
            g.agregar_arista(p[0],link,1)
   return g

def leer_archivos():
   if len(sys.argv) != CANT_PARAMETROS:
        raise ValueError("ERROR: Parámetros inválidos")

   try:
      archivoPaginas = open(sys.argv[ARCHIVO_PAGINAS])
   except:
      raise ValueError(f"ERROR: Lectura de archivos")
   
   return archivoPaginas

def cargar_vertices(archivoPaginas):
   g = Grafo(es_dirigido=True)
   for pagina in archivoPaginas:
      p = pagina.strip().split("\t")
      if p[0]  not in g.obtener_vertices():
         g.agregar_vertice(p[0])
   return g
      
def cargar_aristas(g, archivoPaginas):
