<template>
<label class = "linxreader">
<span class=""> {{  word  }} </span> <!-- TODO assign dynamic id -->
</label>

</template>



<script>
    export default {
        data() {
            return {
                word: ""
            }
        },
        methods: {
            method() {


            }


        },
        beforeMount() {
            this.word = this.word;
        }
        
    }


</script>