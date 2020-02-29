new Vue({
    delimiters: ['${', '}'],
    el: '#tag-create',
    data() {
        return {
            tags: ["hoge", "piyo", "fuga"],
            canEnter: false,
        }
    },
    methods: {
        enter(target) {
            //日本語の確定で EnterキーのKeyupが発生するのを抑止
            if (!this.canEnter) return
            if (typeof target.value === "string" && target.value.trim() !== "") {
                this.tags.push(target.value.trim().replace(/,/, ""))
                target.value = ""
            }
            this.canEnter = false
        },
        remove(index) {
            this.tags.splice(index, 1);
        },
        submitTag: async () => {
            try {
                const result = await axios.post('/post-tag', this.tags)
                window.alert(result)
            } catch (error) {
                window.alert(error)
            }
        }
    },
})
