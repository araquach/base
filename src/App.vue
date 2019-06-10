<template>
    <div>
        <div>
            <section class="hero is-fullheight">
                <div class="hero-body">
                    <div class="container columns is-centered">
                        <home-component v-show="page"></home-component>
<!--                            <info-component v-show="page"></info-component>-->
                    </div>
                </div>
            </section>
        </div>
    </div>
</template>