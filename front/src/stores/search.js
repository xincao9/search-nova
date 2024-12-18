import {ref} from 'vue';
import {defineStore} from 'pinia';

export const useSearchAnswerStore = defineStore('searchAnswer',
  () => {
    const answer = ref([]);
    let loadStatus = ref(false);

    const addAnswer = (item) => {
      const index = answer.value.findIndex(function (record) {
        return record === item;
      });
      if (index >= 0) {
        answer.value.splice(index, 1);
      }
      answer.value.unshift(item);
      save();
    }

    const getAnswer = () => {
      // if (loadStatus.value === false) {
      //   load();
      // }
      load();
      return answer.value || null;
    }

    const load = () => {
      answer.value = JSON.parse(localStorage.getItem("searchAnswer")) || [];
      loadStatus.value = true;
    }

    const save = () => {
      localStorage.setItem("searchAnswer", JSON.stringify(answer.value));
    }
    return {addAnswer, getAnswer};
  })
