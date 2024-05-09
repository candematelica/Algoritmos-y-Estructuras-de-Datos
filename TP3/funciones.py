from grafo import Grafo
from collections import deque
from collections import Counter
import sys
import random

sys.setrecursionlimit(70000)

COEF_AMORTIGUACION = 0.85
MAX_ITERACIONES = 50

def bfs(g, origen):
    visitados = set()
    padres = {}
    orden = {}
    padres[origen] = None
    orden[origen] = 0
    visitados.add(origen)
    q = deque()
    q.append(origen)
    while q:
        v = q.popleft()
        for w in g.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                orden[w] = orden[v] + 1
                visitados.add(w)
                q.append(w)
    return orden, padres

def reconstruir_camino(padres, origen, destino):
    v = destino
    camino = []
    while origen != v:
        if v not in padres:
            break
        camino.append(v)
        v = padres[v]
    if not camino:
        return None
    camino.append(v)
    camino.reverse()
    return camino

def cfcs(g):
    cfcs = []
    visitados = set()
    for v in g.obtener_vertices():
        if v not in visitados:
            dfs_cfc(g, v, visitados, {}, {}, deque(), set(), cfcs, [0])      
    return cfcs

def dfs_cfc(g, v, visitados, orden, mas_bajo, pila, apilados, cfcs, contador):
    orden[v] = mas_bajo[v] = contador[0]
    contador[0] += 1
    visitados.add(v)
    pila.append(v)
    apilados.add(v)
    for w in g.adyacentes(v):
        if w not in visitados:
            dfs_cfc(g, w, visitados, orden, mas_bajo, pila, apilados, cfcs, contador)
        if w in apilados:
            mas_bajo[v] = min(mas_bajo[v], mas_bajo[w])
    
    if orden[v] == mas_bajo[v]:
        nueva_cfc = []
        while True:
            w = pila.pop()
            apilados.remove(w)
            nueva_cfc.append(w)
            if w == v:
                break
        cfcs.append(nueva_cfc)

def orden_topologico(g, vertices):
    g_ent = grados_entrada(g, vertices)
    resultado = []
    q = deque()
    
    for v in vertices:
        if g_ent[v] == 0:
            q.append(v)
    
    while q:
        v = q.popleft()
        if v not in vertices:
            continue
        resultado.append(v)
        for w in g.adyacentes(v):
            g_ent[w] -= 1
            if g_ent[w] == 0:
                q.append(w)
    
    if len(resultado) != len(vertices):
        return None
    
    return resultado

def grados_entrada(g, vertices):
    g_ent = {}
    for v in vertices:
        g_ent[v] = 0
    
    for v in vertices:
        for w in g.adyacentes(v):
            if w not in g_ent:
                g_ent[w] = 0
            g_ent[w] += 1
            
    return g_ent

def vertices_entrada(g):
    v_ent = {}
    
    for v in g.obtener_vertices():
        v_ent[v] = []
    
    for v in g.obtener_vertices():
        for w in g.adyacentes(v):
            v_ent[w].append(v)
            
    return v_ent

def pageranks(g):
    v_ent = vertices_entrada(g)
    pr = {}
    g_longitud = len(g.obtener_vertices())
    
    for v in g.obtener_vertices():
        pr[v] = 1/g_longitud

    for v in pr:
        suma = 0
        for w in v_ent[v]:
            suma += (pr[w] / len(g.adyacentes(w)))
            
        pr[v] = (1 - COEF_AMORTIGUACION)/g_longitud + COEF_AMORTIGUACION * suma
        
    return pr

def label_propagation(g):
    v_ent = vertices_entrada(g)
    vertices = g.obtener_vertices()
    random.shuffle(vertices)
    label = {}
    
    for i, v in enumerate(vertices):
        label[v] = i
    
    i = 0
    while i < MAX_ITERACIONES:
        cambios = 0
        for v in label:
            lo_tienen_de_ady = [label[w] for w in v_ent[v]]
            contador = Counter(lo_tienen_de_ady)
            if contador:
                nueva_label = contador.most_common(1)[0][0]
                if label[v] != nueva_label:
                    label[v] = nueva_label
                    cambios += 1
        if cambios == 0:
            break
        
        i += 1
            
    return label
