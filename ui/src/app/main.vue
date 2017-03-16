
<style>
body {
  padding-top: 40px;
}
</style>

<template>
<div>
  <div class="jumbotron">
    <h1>Hello, {{ who }}.</h1>
    <p>Get the connection graph of each bitmarkd networks</p>
  </div>  



  <conn-graph></conn-graph>
</div>

</template>

<script>

import connGraph from './connGraph.vue'

export default {
  components: {
    "conn-graph": connGraph
  },

  data () {
    return {
      "who": "world"
    }
  }
}
</script>
