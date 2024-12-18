export default {
  mounted(el, binding) {
    const {value} = binding;
    if (!value || !Array.isArray(value) || value.length < 2) {
      console.warn('v-highlight needs an array with at least two elements: [searchQuery, highlightClass]');
      return;
    }

    const [searchQuery, highlightClass] = value;
    const regex = new RegExp(`(${searchQuery})`, 'gi');

    // Create a span for each match and apply the highlight class
    el.innerHTML = el.innerHTML.replace(regex, `<span class="${highlightClass}">$1</span>`);
  },
  updated(el, binding) {
    this.mounted(el, binding);
  }
};
