<template>
<label class = "linxreader">
<!-- TODO assign dynamic id -->
<button @click="OnClick">  {{ renderWord }}  </button>
</label>
    
</template>



<script>
    export default {
        props: ['renderWord'],
        data() {
            return {
                word: ""
            }
        },
        methods: {

                OnClick() {
                    
                    this.$emit("linxreaderload",word);


                }


        },
        beforeMount() {
//            this.word = this.word;
        }
        
    }


</script>