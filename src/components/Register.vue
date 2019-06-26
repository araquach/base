<template>
    <div class="section form-copy">
        <h1 class="is-hidden">
            Base Hairdressing
        </h1>
        <div class="columns is-centered">
            <div class="column is-6 has-text-centered">
                <a href="/">
                    <figure class="image">
                        <img src="dist/images/base_logo.svg" alt="Base Hairdressing">
                    </figure>
                </a>
            </div>
        </div>
        <h2 class="subtitle is-3 has-text-white has-text-centered">Register your interest with us!</h2>
        <p>If Base seems like the perfect salon environment for you, then fill in the form and we'll be in touch soon! </p>
        <form @submit="checkForm" action="/register" method="post">

            <div v-if="errors.length" class="box has-text-danger">
                <p><strong>Please correct the following:</strong></p>
                <ul>
                    <li v-for="error in errors">{{ error }}</li>
                </ul>
            </div>

            <div class="field">
                <label class="label">Name</label>
                <div class="control">
                    <input class="input" v-model="name" name="name" type="text" placeholder="Your Name">
                </div>
            </div>
            <div class="field">
                <label class="label">Mobile Number</label>
                <div class="control">
                    <input class="input" v-model="mobile" name="mobile" type="text" placeholder="Your Mobile Number">
                </div>
            </div>
            <div class="field">
                <label class="label">Current Position</label>
                <div class="control">
                    <div class="select">
                        <select v-model="position" name="position">
                            <option value="default">Please select</option>
                            <option value="employed">Employed Stylist</option>
                            <option value="chair renter">Chair Renter</option>
                            <option value="mobile">Mobile Hairdresser</option>
                            <option value="other">Other</option>
                        </select>
                    </div>
                </div>
            </div>
            <br>
            <div class="field">
                <div class="control">
                    <button class="button is-primary" type="submit" value="submit">Submit</button>
                </div>
            </div>
        </form>
    </div>
</template>

<script>

    export default {

        data() {
            return {
                errors: [],
                name: null,
                mobile: null,
                position: null
            }
        },

        methods:{
            checkForm: function (e) {
                this.errors = [];

                if (!this.name) {
                    this.errors.push('Name required.');
                }
                if (!this.mobile) {
                    this.errors.push('Mobile Number required.');
                } else if (!this.validMobile(this.mobile)) {
                    this.errors.push('Valid Mobile Number required.');
                }
                if (!this.position) {
                    this.errors.push('Position required')
                }

                if (!this.errors.length) {
                    return true;
                }

                e.preventDefault();
            },

            validMobile: function (mobile) {
                var re = /^((\+44\s?|0)7([45789]\d{2}|624)\s?\d{3}\s?\d{3})$/
                return re.test(mobile);
            }
        }

    }
</script>