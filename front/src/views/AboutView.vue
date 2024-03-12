<template>
  <div >
    <div>
      <h1><b>Exercício 1</b></h1>
      <button @click="fetchData">Clique aqui</button>
    </div>
    <div id="resultado">{{ resposta }}</div>
    <br>
    <br>
    <br>
  <div>
      <h1><b>Exercício 2</b></h1>
      <button @click="fetchData2">Clique aqui</button>
    </div>
    <div id="resultado2">{{ resposta2 }}</div>
    <br>
    <br>
    <br>
  <div>
      <h1><b>Exercício 3</b></h1>
      <button @click="fetchData3">Clique aqui</button>
    </div>
    <div id="resultado3">{{ resposta3 }}</div>
  </div>

  
</template>

<script>
import api from '../service/index.js';

//requisição
export default {
  data() {
    return {
      resposta: '',
      resposta2: '',
      resposta3: ''
    };
  },
  methods: {
    async fetchData() {
      try {
        const response = await api.get('/teste');
        this.resposta = response.data;
      } catch (error) {
        console.error('Erro ao fazer requisição:', error);
      }
    },
    async fetchData2() {
      try {
        const response = await api.get('/teste/2001');
        this.resposta2 = response.data;
      } catch (error) {
        console.error('Erro ao fazer requisição:', error);
      }
    },
    async fetchData3() {
      try {
        const response = await api.get('/exercicio3?valor=10&quantidade=3');
        this.resposta3 = response.data;
      } catch (error) {
        console.error('Erro ao fazer requisição:', error);
      }
    }
  }
};

</script>





<style scoped>
body {
  background-color: #000;
  /* Cor preta */
  color: #fff;
  /* Cor branca para o texto */
  text-align: center;
  /* Centraliza o conteúdo */
}

h1 {
  margin-bottom: 20px;
  /* Espaço abaixo do título */
}

button {
  font-size: 24px;
  /* Tamanho da fonte do botão */
  padding: 10px 20px;
  /* Espaçamento interno do botão */
}
</style>
