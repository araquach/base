<template>
    <div>
        <div>
            <section class="hero is-fullheight">
                <div class="hero-body">
                    <div class="container columns is-centered">
                        <div class="column is-5">
                            <home-component v-show="page"></home-component>
                            <info-component v-show="page"></info-component>
                        </div>
                    </div>
                </div>
            </section>
        </div>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                page: 'home'
            }

        }
    }
</script>