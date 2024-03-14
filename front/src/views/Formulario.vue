<template>
  <div class="form-container">
    <form @submit.prevent="submitForm" class="form">
      <h1 class="form-title">Formulário WEB 2</h1>

      <div class="form-group">
        <label for="texto">Nome:</label>
        <input type="text" id="texto" v-model="form.texto" required>
      </div>

      <div class="form-group">
        <label for="inteiro">Insira um número:</label>
        <input type="number" id="inteiro" v-model.number="form.inteiro" required>
      </div>

      <div class="form-group">
        <input type="checkbox" id="booleano" v-model="form.booleano">
        <label for="booleano">Aceita os termos</label>
      </div>

      <div class="form-group">
        <label for="opcaoSelect">Seleciona uma Opção:</label>
        <select id="opcaoSelect" v-model="form.opcaoSelect" required>
          <option name="check" value="opcao1">Opção 1</option>
          <option name="check" value="opcao2">Opção 2</option>
          <option name="check" value="opcao3">Opção 3</option>
        </select>
      </div>

      <div class="form-group">
        <label>Escolha uma Opção:</label>
        <div class="radio-group">
          <input type="radio" id="radio1" value="opcao1" v-model="form.opcaoRadio">
          <label for="radio1">Opção 1</label>
          <input type="radio" id="radio2" value="opcao2" v-model="form.opcaoRadio">
          <label for="radio2">Opção 2</label>
          <input type="radio" id="radio3" value="opcao3" v-model="form.opcaoRadio">
          <label for="radio3">Opção 3</label>
        </div>
      </div>

      <button type="submit" class="submit-button">Enviar</button>
    </form>

    <!-- Seção de resultados -->
    <div class="result-container" v-if="showResults">
      <h2>Resultados:</h2>
      <p>Nome: {{ form.texto }}</p>
      <p>Número: {{ form.inteiro }}</p>
      <p>Aceita os termos: {{ form.booleano ? 'Sim' : 'Não' }}</p>
      <p>Opção selecionada: {{ form.opcaoSelect }}</p>
      <p>Opção de rádio selecionada: {{ form.opcaoRadio }}</p>
    </div>
  </div>
</template>

<script>
import api from '../service';
export default {
  data() {
    return {
      form: {
        texto: '',
        inteiro: null,
        booleano: false,
        opcaoSelect: '',
        opcaoRadio: ''
      },
      showResults: false // Para controlar a exibição dos resultados
    };
  },
  methods: {
    async submitForm() {
      await api.post('/formulario', this.form).then(response => {
        console.log(response.data);
        this.showResults = true; // Mostrar os resultados após o envio do formulário
      }).catch(error => {
        console.error('Erro:', error);
      });
    }
  }
}
</script>

<style>
/* Estilos do formulário */
.form-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.form {
  max-width: 400px;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

.form-title {
  text-align: center;
  margin-bottom: 20px;
}

.form-group {
  margin-bottom: 15px;
}

label {
  display: block;
  margin-bottom: 5px;
}

input[type="text"],
input[type="number"],
select {
  width: 100%;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

.radio-group {
  display: flex;
  align-items: center;
}

input[type="radio"] {
  margin-right: 5px;
}

.submit-button {
  width: 100%;
  padding: 10px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.submit-button:hover {
  background-color: #0056b3;
}

/* Estilos dos resultados */
.result-container {
  margin-top: 20px;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
}
</style>
