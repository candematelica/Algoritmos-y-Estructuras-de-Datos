import sys
from grafo import Grafo
from grafo import chequear_pertenencia
from funciones import *
import heapq

sys.setrecursionlimit(70000)

COMANDO_LISTAR_OPERACIONES = "listar_operaciones"
COMANDO_CAMINO = "camino"
COMANDO_MAS_IMPORTANTES = "mas_importantes"
COMANDO_CONECTIVIDAD = "conectados"
COMANDO_LECTURA = "lectura"
COMANDO_DIAMETRO = "diametro"
COMANDO_RANGO = "rango"
COMANDO_COMUNIDAD = "comunidad"
COMANDO_NAVEGACION = "navegacion"
MAX_NAVEGACION = 20

def leer_comandos(g: Grafo()):
    g_cfcs = cfcs(g)
    pr = pageranks(g)
    lp = label_propagation(g)
    
    diam_calc = False
    camino = []
    diam = 0
    
    for linea_comando in sys.stdin:
        comando = linea_comando.strip().split(" ", 1)
        if comando[0] == COMANDO_LISTAR_OPERACIONES:
            listar_operaciones()
        elif comando[0] == COMANDO_CAMINO:
            origen, destino = comando[1].split(",")
            camino_min = camino_mas_corto(g, origen, destino)
            if camino_min is not None:
                print(" -> ".join(map(str, camino_min)))
                print(f"Costo: {len(camino_min)-1}")
            else:
                print("No se encontro recorrido")
        elif comando[0] == COMANDO_MAS_IMPORTANTES:
            cant_paginas = comando[1]
            importantes = mas_importantes(int(cant_paginas), pr)
            print(", ".join(map(str, importantes)))
        elif comando[0] == COMANDO_CONECTIVIDAD:
            pagina = comando[1]
            conectados = conectividad(g, pagina, g_cfcs)
            print(", ".join(map(str, conectados)))
        elif comando[0] == COMANDO_LECTURA:
            paginas = comando[1].split(",")
            orden = lectura(g, paginas)
            if orden is not None:
                print(", ".join(map(str, orden)))
            else:
                print("No existe forma de leer las paginas en orden")
        elif comando[0] == COMANDO_DIAMETRO:
            if not diam_calc:
                diam_calc = True
                camino, diam = diametro(g)
            print(" -> ".join(map(str, camino)))
            print(f"Costo: {diam}")
        elif comando[0] == COMANDO_RANGO:
            origen, rango = comando[1].split(",")
            contador = paginas_en_rango(g, origen, int(rango))
            print(contador)
        elif comando[0] == COMANDO_COMUNIDAD:
            pagina = comando[1]
            comu = comunidad(g, pagina, lp)
            print(", ".join(map(str, comu)))
        elif comando[0] == COMANDO_NAVEGACION:
            origen = comando[1]
            resultado_nav = navegacion_por_primer_link(g, origen)
            print(" -> ".join(map(str, resultado_nav)))
        else:
            print(f"ERROR: Comando invalido", comando[0])

def listar_operaciones():
    print(COMANDO_CAMINO)
    print(COMANDO_MAS_IMPORTANTES)
    print(COMANDO_CONECTIVIDAD)
    print(COMANDO_LECTURA)
    print(COMANDO_DIAMETRO)
    print(COMANDO_RANGO)
    print(COMANDO_COMUNIDAD)
    print(COMANDO_NAVEGACION)
    
def camino_mas_corto(g, origen, destino):
    chequear_pertenencia([origen, destino], g.obtener_vertices())
    _, padres = bfs(g, origen)
    return reconstruir_camino(padres, origen, destino)
    
def mas_importantes(n, pr):
    paginas =[]
    for p in pr:
        heapq.heappush(paginas, (pr[p], p))
    
    importantes = heapq.nlargest(n, paginas)
        
    return [p[1] for p in importantes]

def conectividad(g, pagina, g_cfcs):
    chequear_pertenencia([pagina], g.obtener_vertices())
    for cfc in g_cfcs:
        if pagina in cfc:
            return cfc

def lectura(g, paginas):
    chequear_pertenencia(paginas, g.obtener_vertices())
    orden = orden_topologico(g, paginas)
    if orden is not None:
        orden.reverse()
    return orden

def diametro(g):
    mayor_cm = []
    mayor_long_cm = 0

    for v in g.obtener_vertices():
        orden, padres = bfs(g, v)
        for w in g.obtener_vertices():
            if w not in orden:
                continue
            dist = orden[w]
            if dist > mayor_long_cm:
                camino = reconstruir_camino(padres, v, w)
                if camino is not None:
                    mayor_long_cm = dist
                    mayor_cm = camino

    return mayor_cm, mayor_long_cm
    
def paginas_en_rango(g, origen, rango):
    chequear_pertenencia([origen], g.obtener_vertices())
    contador = 0
    orden, _ = bfs(g, origen)
    for v in g.obtener_vertices():
        if orden[v] == rango:
            contador += 1
        
    return contador
    
def comunidad(g, pagina, lp):
    chequear_pertenencia([pagina], g.obtener_vertices())
    comu = []
    for v in lp:
        if lp[v] == lp[pagina]:
            comu.append(v)
    
    return comu
    
def navegacion_por_primer_link(g, origen):
    chequear_pertenencia([origen], g.obtener_vertices())
    i = 0
    resultado = []
    actual = origen
    while i != MAX_NAVEGACION or len(g.adyacentes(actual)) > 0:
        resultado.append(actual)
        ady = g.adyacentes(actual)
        if not ady:
            break
        actual = ady[0]
        i += 1
        
    return resultado
