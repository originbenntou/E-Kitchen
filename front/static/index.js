const VeeValidate = window.VeeValidate
const ValidationProvider = VeeValidate.ValidationProvider
const ValidationObserver = VeeValidate.ValidationObserver
const VeeValidateRules = window.VeeValidateRules
const required = VeeValidateRules.required
const email = VeeValidateRules.email
const password = VeeValidateRules.password

required.message = '必須項目です'
email.message = 'メールの形式ではありません'

VeeValidate.extend('required', required)
VeeValidate.extend('email', email)

Vue.component('ValidationProvider', ValidationProvider)
Vue.component('ValidationObserver', ValidationObserver)

new Vue({
    delimiters: ['${', '}'],
    el: '#app',
    data: {
        email: '',
        password: '',
    },
    methods: {
        submit: () => {
            alert('ログイン成功！')
        }
    }
})
