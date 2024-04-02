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
        <input type="checkbox" id="booleano" v-model="form.booleano" required>
        <label for="booleano">Aceita os termos</label>
      </div>

      <div class="form-group">
        <label for="opcaoSelect">Seleciona uma Opção:</label>
        <select id="opcaoSelect" v-model="form.opcaoSelect" required>
          <option value="" disabled selected>Selecione uma opção</option>
          <option value="opcao1">Opção 1</option>
          <option value="opcao2">Opção 2</option>
          <option value="opcao3">Opção 3</option>
        </select>
      </div>

      <div class="form-group">
        <label>Escolha uma Opção:</label>
        <div class="radio-group" >
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

   
    <div class="result-container" v-if="showResults">
      <h2>Resultados:</h2>
      <p><span class="text-style">Nome: </span>{{ tabela1.texto||""}}</p>
      <p><span class="text-style">Número: </span>{{ tabela1.inteiro||0 }}</p>
      <p><span class="text-style">Aceita os termos:</span> {{ tabela1.booleano ? 'Sim' : 'Não' }}</p>
      <p><span class="text-style">Opção selecionada:</span> {{ tabela1.opcaoSelect }}</p>
      <p><span class="text-style">Opção de rádio selecionada: </span>{{ tabela1.opcaoRadio }}</p>
    </div>
  </div>
</template>

<script>
import api from '../service';
export default {
  data() {
    return {
      tabela1: {
        texto: '',
        inteiro: null,
        booleano: false,
        opcaoSelect: '',
        opcaoRadio: ''
      },
      form: {
        texto: '',
        inteiro: null,
        booleano: false,
        opcaoSelect: '',
        opcaoRadio: ''
      },
      showResults: false 
    };
  },
  methods: {

    validate() {
      let flag = true, errorMensage = "";
      if (!this.name_validade()) {
        errorMensage += "O nome precisa ter no mínimo 2 e no máximo 255 caracteres\n"
        flag = false
      }

      if (!this.number_validade()) {
        errorMensage += "O número precisa ser maior que 0 e menor que 1000\n"
        flag = false
      }

      if (!this.radio_validate()) {
        errorMensage += "Selecione uma opção\n"
        flag = false
      }

      return [flag, errorMensage]
    },

    name_validade() {
      let size = this.form.texto.length
      if (size >= 2 && size <= 255) return true;
      return false;
    },

    number_validade() {
      let num = this.form.inteiro
      if (num < 0 || num >= 1000) return false;
      return true;

    },

    radio_validate() {
      let op = this.form.opcaoRadio;

      if (op == '') return false;
      return true;
    },

    reset_form(){
      this.form.showResults = true;
      this.form.texto = '',
      this.form.inteiro = null,
      this.form.booleano = false,
      this.form.opcaoSelect = '',
      this.form.opcaoRadio = ''      
    },

    async submitForm() {
      try {
        let [flag, errorMensage] = this.validate();
        if(flag) {
          const response = await api.post(api.defaults.baseURL + '/formulario', this.form);
          console.log(response.data);
          this.tabela1 = response.data
          this.showResults = true;
          this.reset_form();

        } else alert(errorMensage)
      } catch (error) {
        console.error('Erro:', error);
      }
    }
  }
}
</script>

<style>
.form-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.form {
  width: 400px;
  padding: 30px;
  border-radius: 10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  background-color: #bfb3e9;
}

.form-title {
  text-align: center;
  margin-bottom: 20px;
  color: #333;
}

.form-group {
  margin-bottom: 20px;
}

label {
  display: block;
  margin-bottom: 5px;
  color: #060606;
}

input[type="text"],
input[type="number"],
select {
  width: 100%;
  padding: 10px;
  border: 1px solid #181414;
  border-radius: 5px;
  font-size: 16px;
}

.checkbox-group {
  display: flex;
  align-items: center;
}

.checkbox-group label {
  margin-left: 5px;
}

.radio-group {
  display: flex;
  align-items: normal;
}

input[type="radio"] {
  margin-right: 10px;
}

.submit-button {
  width: 100%;
  padding: 12px;
  background-color: #007bff;
  color: #151414;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.submit-button:hover {
  background-color: #0056b3;
}

.result-container {
  margin-top: 20px;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 0 10px rgb(27, 20, 20);
  background-color: #7c66cc;
  color: #111111d5;
}

.result-container h2 {
  text-align: center;
  margin-bottom: 10px;
  color: #000000;
  padding-right: 100px;
  white-space: nowrap;
  
}

.text-style{
  font-size: 15px;
  font-weight: bold;
}

.radio-group label{
  padding-right: 50px;
  white-space: nowrap;
}
</style>
