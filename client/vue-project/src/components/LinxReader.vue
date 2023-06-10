<template>
<label class = "linxreader">
<!-- TODO assign dynamic id -->
<button 
@click="$emit('stateChanged',learnedReturn)"
>
{{ renderWord }} </button>
</label>
    
</template>



<script>
    export default {
        props: {renderWord : String,
                wordId : String,
                learned : Number,
                created() {
                    this.wordReturn = this.renderWord;
                    this.idReturn = this.wordId;
                    this.learnedReturn = this.learned;
                }
            },
        data() {
            return {
                wordReturn: "",
                idReturn: null,
                learnedReturn: 0,
                sentence : {
            sentenceitem : {
              unknownword : {
                //Colors, rendering stuff
              },
              knownword : {
                tha : String
                //Colors, rendering stuff
              }
            }
          }
            }
        },
        methods: {

                OnClick() {
                    
                    

                }


        },
        mounted() {
//            this.word = this.word;

        }
        
    }


</script>

<style>


</style>