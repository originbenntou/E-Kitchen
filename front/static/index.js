const VeeValidate = window.VeeValidate
const ValidationProvider = VeeValidate.ValidationProvider
const ValidationObserver = VeeValidate.ValidationObserver
const VeeValidateRules = window.VeeValidateRules
const required = VeeValidateRules.required
const email = VeeValidateRules.email

required.message = '必須項目です'
email.message = 'メールアドレスの形式ではありません'

VeeValidate.extend('required', required)
VeeValidate.extend('email', email)

// custom rule
const password = {
    message: '半角英数字記号7~33桁ではありません',
    validate(value) {
        const reg = new RegExp(/^([a-zA-Z0-9!-/:-@¥[-`{-~]{7,33})$/);
        return reg.test(value)
    }
};
VeeValidate.extend('password', password);

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
            // alert('ログイン成功！')
        },
    }
})
