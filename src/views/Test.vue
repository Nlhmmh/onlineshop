<template>
  <div>
    <!-- Expression, Mustache Form, Interpolation -->
    <h1>Hello</h1>
    <h1>{{ name }}</h1>

    <!-- v-form -->
    <h3>Form</h3>
    FirstName : <input type="text" v-model="firstName" />
    <h5>{{ firstName }}</h5>

    <!-- v-for -->
    <div class="my-row">
      <div v-for="(item, index) in itemList" :key="index" class="my-card">
        <p>{{ item.name }}</p>
        <p>{{ item.count }}</p>
      </div>
    </div>

    <!-- v-bind -->
    <div v-bind:class="testClass">Test for v-bind</div>
    <button :disabled="isBtnDisabled">Confirm</button>

    <!-- v-on -->
    <button @click="onClickEvent">Click Me!</button>
    <h4>{{ onClickData }}</h4>
  </div>
</template>

<script>
export default {
  data() {
    return {
      name: "NLH",
      firstName: "",
      itemList: [
        { name: 'pizza is "OK"', count: 3 },
        { name: "cake", count: 5 },
        { name: "coffee", count: 4 },
      ],
      testClass: "myDynamicClassBlue",
      isBtnDisabled: false,
      onClickData: true,
    };
  },
  methods: {
    onClickEvent: function () {
      console.log("Clicked!");
      this.onClickData = !this.onClickData;
    },
  },
};
</script>

<style scoped>
.my-row {
  display: flex;
  justify-content: center;
}

.my-card {
  border: 1px solid black;
  width: 100px;
  height: 100px;
  margin: 10px;
  padding: 10px;
}
.myDynamicClassRed {
  background: red;
  width: 100px;
  height: 100px;
}
.myDynamicClassBlue {
  background: blue;
  width: 100px;
  height: 100px;
}
</style>
