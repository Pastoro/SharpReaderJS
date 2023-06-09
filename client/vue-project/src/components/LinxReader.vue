<template>
<label class = "linxreader">
<!-- TODO assign dynamic id -->
<button >  {{ renderWord}}  </button>
</label>
    
</template>



<script>
    export default {
        props: {renderWord : String,
                wordId : String,
                learned : Object
            },
        data() {
            return {
                word: "",
                id: null
            }
        },
        methods: {

                OnClick() {
                    
                    this.$emit("linxreaderload",this.word);


                }


        },
        mounted() {
//            this.word = this.word;

        }
        
    }


</script>

<style>


</style>