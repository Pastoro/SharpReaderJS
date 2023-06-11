<template>
    <div id="linx">
        
        <br>
      <h1>This is the Linx Page</h1>
      <linximporter @load="LoadHandler"></linximporter> <!-- TODO assign dynamic id -->
        <div class="grid-container"  > <!-- v-for state === unknown word, class = unknownword  -->
          <!-- <div :style="{color : activeColor}" :id="`container-${index}`">{{ word }}   XXX possibly add a check in here to change the colour accordingly-->
          <linxreader    
          v-for="element in words"
          :renderWord="element.word" 
          :learned="element.state"
          :class="linxreadermethod(element.word,element.state)"
          @click="element.state = ClickEvent(element.state)"
          >
            
        
          </linxreader>
        </div>
    
      </div>


  </template>
<script> //XXX you can probably use CSS to move the menu component around whilst only having it be bound to the grid-container class
    import linximporter from "../components/Linx.vue"
    import linxreader from "../components/LinxReader.vue"
    import axios from 'axios';
    import LinxMenu from "../components/LinxMenu.vue";


    export default{
      mounted : function() {
          this.HandleTextRender();
        },
        data () {

          return {
           // renderWord: '',
           // learned: '',
          classValue: "",
          activeColor: 'red',
          mywordState : null,
          rawhtml : '<span style="color: red">This should be red.</span>',
          wordMap : Map,
          separateState : {},
          words: [
            //arrayword:"", arraystate:{}
          ]
          }
          
        },
        components: {
    linximporter,
    linxreader,
    LinxMenu,
    LinxMenu,
    LinxMenu
}
        ,
        methods: {
          LoadHandler(textMap, wordState) {
            //this.wordMap = textMap;
            this.mywordState = wordState;
            console.log(this.mywordState);
            var word;
            var state;

            for (const [key, value] of textMap) {
              word = key;
              state = value;
              this.words.push({word, state});
            }
            console.log(this.words);
            //this.words.forEach(word => {console.log(word)});
          },
          HandleTextRender(){

          axios({ method: "GET", params: {QueryWord: "prima"}, responseType: 'json', url: " http://localhost:8030/fetchWord"}).then(function (response) {
            console.log(response.data);
            //console.log(JSON.stringify(JSON.parse(response), null, 2));
          });
      


        },
        LinxReaderLoaded : function(event) {
          console.log(event.target.id);
          //Make call when this is called to call back to server to check for user data and dynamically update the CSS styling from there

        },
        linxreadermethod: function(renderWord, state) {
            var classA = '';
            console.log(renderWord);
            console.log(state);
            if (state === this.mywordState.DefaultState) {
              classA = 'wordthingUnknown';
            }
            else if (state === this.mywordState.LookedUpState) {
              classA = 'wordthingKnown';
            }
            return classA;
          },
          ClickEvent: function(state) {
            console.log(state);
            state = this.mywordState.LookedUpState;
            return state;
          }

      },
        computed: {

        }

        
    }
  
  

</script>

<style>
.grid-container {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  background-color: #10436d;
  padding: 10px;
  grid-auto-rows: minmax(100px, auto);
}
.wordthingKnown {
  display:inline;
  background-color: rgba(255, 0, 34, 0.8);
  border: 1px solid rgba(25, 0, 255, 0.8);
  padding: 20px;
  font-size: 30px;
  text-align: center;
  font: bold;
}
.wordthingUnknown {
  display:inline;
  background-color: rgba(0, 255, 179, 0.8);
  border: 1px solid rgba(0, 0, 0, 0.8);
  padding: 20px;
  font-size: 30px;
  text-align: center;
}
.grid-container linxreader{
  position: relative;
}

</style>