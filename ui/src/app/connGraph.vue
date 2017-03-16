
<style type="text/css" scoped>
.flex {
  display: flex;
}

.graph {
  width: 480px;
}

</style>

<template>
<div>
  <div class="panel panel-default">
    <div class="panel-body">
      <div class="flex">
        <div style="width: 40%; border-right: solid 1px #e4e4e4; padding-right: 5px; margin-right: 5px">
          <form class="form-horizontal">
            <div class="form-group">
              <label for="textArea" class="col-md-2 control-label">Nodes</label>
              <div class="col-md-10">
                <textarea class="form-control" rows="3" id="textArea"></textarea>
                <span class="help-block">Node addresses, split by comma.</span>
              </div>
            </div>
            <div class="form-group">
              <div class="col-md-10 col-md-offset-2">
                <button type="button" class="btn btn-primary">Draw</button>
              </div>
            </div>
          </form>
        </div>
        <div>
          <img class="graph" :src="pngData">
        </div>
      </div>
    </div>
  </div>
</div>
</template>

<script type="text/javascript">

import axios from 'axios';  

export default {

  methods: {
    getImages () {
      axios({
        method: 'post',
        url: "/v1/graph/",
        data: {}
      }).then((response) => {
        console.log('update gif')
        this.pngData = "data:image/gif;base64," + response.data
      })
      .catch((response) => {
        
      })
    }
  },

  mounted () {
    this.getImages()
    setInterval(()=> this.getImages(), 600000)
  },

  data () {
    return {
      pngData: "file:///Users/jimyeh/.go/src/github.com/lemonlatte/conn2graphviz/connections.gif"
    }
  }
}
</script>  