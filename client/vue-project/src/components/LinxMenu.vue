
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
        async getDictionaryResult(term){ //TODO supposedly computed properties are better and watch is discouraged
            //TODO Leave this as is for now, assuming we're gonna' use a third-party dictionary for this anyways.
            console.log(term);
            console.log("FIRED");

            var returnType = ((await axios({ method: "GET", params: {QueryWord: term}, responseType: 'json', url: " http://localhost:8030/fetchWord"})).data);
          console.log(returnType);
          this.$emit("gotDictionaryResult",returnType);
          return returnType;
        }
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