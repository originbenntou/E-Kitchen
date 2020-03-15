new Vue({
    delimiters: ['${', '}'],
    el: '#tag-create',
    data: {
        tags: [],
        canEnter: false,
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
        submitTag: async function() {
            let params = new URLSearchParams();
            params.append('data', JSON.stringify(this.tags))
            try {
                const result = await axios.post('/tag-register', params, {
                    'Content-Type': 'application/x-www-form-urlencoded',
                })
            } catch (error) {
                window.alert("error:" + error)
            }
        }
    },
})
