
<template>
    <label class ="linxmenu">

        

        <div  class="linxmenuPopup" >

            <p class="search" > {{ searchedTerm }} </p>
            <p class="query"> {{ queryResult }}</p>

        </div>



        
    </label>




</template>

<script>
import axios from 'axios';


export default{
    
    data(){
        return{
            searchTerm: "",
            resultTerm: "",
            isMounted: false
        }
    },
    props: {
        searchedTerm: String,
        queryResult: String,
        mountedBool: Boolean,
        created(){
            this.searchTerm = this.searchedTerm;
            this.resultTerm = this.queryResult;
        }
    },

    methods:{
        
    },
    mounted(){
        this.$emit('menumounted', true);
        console.log("MOUNTED MENU HERE");
    },
}



</script>



<style>

.linxmenuPopup {
    color: violet;
    background-color: whitesmoke;
    width: 400px;
    height: 500px;
    z-index: 30;
    position: absolute;
}
.query{
    color:maroon;
}


</style>