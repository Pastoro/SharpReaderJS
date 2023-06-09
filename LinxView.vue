<template>
    <div id="linx">
        
        <br>
      <h1>This is the Linx Page</h1>
      <linximporter @load="LoadHandler"></linximporter>
      <div class="content" v-for="[word, state], index in wordMap"> <!-- v-for state === unknown word, class = unknownword  -->
      <div :style="{color : activeColor}" :id="`container-${index}`"> {{ word }}  </div> <!--XXX possibly add a check in here to change the colour accordingly-->
    </div>
    </div>

  </template>
 //XXX Might be better to iterate over the map in the method beforehand
<script>
    import linximporter from "../components/Linx.vue"
    import linxreader from "../components/LinxReader.vue"
    export default{
      mounted : function() {
          this.HandleTextRender();
        },
        data () {
          
          //XXX they've probably done it by simply making different classes for every kind of interaction and then simply switching the interaction on event

          return {
          activeColor: 'red',
          mySentence : {},
          rawhtml : '<span style="color: red">This should be red.</span>',
          wordMap : Map,
          word : "",
          }//one sentence class holds multiple sentenceitems
          //TODO keeping wordState in data return might not be necessary. Keep an array of both the word and the wordState in data return and simply modify the state within a method to allow for freezing. It should still update the rendering when one of the elements of the array changes.
          
        },
        components: {
        linximporter,
        linxreader
        },
        methods: {
          LoadHandler(textMap, sentence) {
            this.wordMap = textMap;
            this.mySentence = sentence;
            this.HandleTextRender();
          },
          HandleTextRender(){

          }


        }

        
    }

</script>

<style>
.html {
  color: red;
}


</style>