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
        <form action="/register">
            <div class="field">
                <label class="label">Name</label>
                <div class="control">
                    <input class="input" type="text" placeholder="Your Name">
                </div>
            </div>
            <div class="field">
                <label class="label">Mobile Number</label>
                <div class="control">
                    <input class="input" type="text" placeholder="Your Mobile Number">
                </div>
            </div>
            <div class="field">
                <label class="label">Current Position</label>
                <div class="control">
                    <div class="select">
                        <select>
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
                    <button class="button is-primary">Submit</button>
                </div>
            </div>
        </form>
    </div>
</template>