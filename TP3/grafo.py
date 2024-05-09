import random

class Grafo:
    def __init__(self, es_dirigido=False, vertices=[]):
        self.es_dirigido = es_dirigido
        self.vertices = {}
        for vertice in vertices:
            self.vertices[vertice] = {}
        
    def agregar_vertice(self, v):
        if v in self.vertices:
            raise ValueError(f"ERROR: El vertice {v} ya pertenece al grafo")
        self.vertices[v] = {}
        
    def borrar_vertice(self, v):
        chequear_pertenencia([v], self.vertices)
        del self.vertices[v]
        for vertice in self.vertices:
            if v in self.vertices[vertice]:
                del self.vertices[vertice][v]
        
    def agregar_arista(self, v, w, peso):
        chequear_pertenencia([v, w], self.vertices)
        self.vertices[v][w] = peso
        if not self.es_dirigido:
            self.vertices[w][v] = peso
        
    def borrar_arista(self, v, w):
        chequear_pertenencia([v, w], self.vertices)
        if w not in self.vertices[v]:
            raise ValueError(f"ERROR: No hay una arista de {v} a {w} en el grafo")
        del self.vertices[v][w]
        if not self.es_dirigido:
            del self.vertices[w][v]
        
    def estan_unidos(self, v, w):
        chequear_pertenencia([v, w], self.vertices)
        if w in self.vertices[v]:
            return True
        return False
        
    def peso_arista(self, v, w):
        chequear_pertenencia([v, w], self.vertices)
        if not w in self.vertices[v]:
            raise ValueError(f"ERROR: No hay una arista de {v} a {w} en el grafo")
        return self.vertices[v][w]
    
    def obtener_vertices(self):
        return list(self.vertices.keys())
        
    def adyacentes(self, v):
        return list(self.vertices[v].keys())
    
    def vertice_aleatorio(self):
        return random.choice(list(self.vertices.keys()))
        
    def __str__(self):
        return str(self.vertices.keys())
    
def chequear_pertenencia(vertices, verticesGrafo):
    for v in vertices:
        if v not in verticesGrafo:
            raise ValueError(f"ERROR: El vertice {v} no pertenece al grafo")
