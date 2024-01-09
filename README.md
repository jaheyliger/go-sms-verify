# go-sms-verify
OTP app using Twilio for SMS verification
 
## Routes

Running commands with npm `npm run [command]`

| route           | description                              |
| :-------------- | :--------------------------------------- |
| `/otp`          | POST: {
    "phoneNumber": "+1123456789"
} |
| `/veriftOPT`    | POST: {
    "user": {
        "phoneNumber": "+1123456789"
    },
    "code": "095993"
}  |
