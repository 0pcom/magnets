## Snipcart payment integration

https://docs.snipcart.com/v3/setup/installation

Include default stylesheet
```
<link rel="stylesheet" href="https://cdn.snipcart.com/themes/v3.1.0/default/snipcart.css" />
```
Include Snipcart

Here, you can decide where you want Snipcart's JS shopping cart to be located in your website.
```
<script async src="https://cdn.snipcart.com/themes/v3.1.0/default/snipcart.js"></script>
<div hidden id="snipcart" data-api-key="YOUR_PUBLIC_API_KEY"></div>
```

an example of a basic setup with Snipcart installed properly.
```
<!DOCTYPE html>
<html>
  <head>
    <title>Hello world</title>
    <link rel="stylesheet" href="https://cdn.snipcart.com/themes/v3.1.0/default/snipcart.css" />
  </head>
  <body>
    <div class="content">
      Your site content here
    </div>

    <script async src="https://cdn.snipcart.com/themes/v3.1.0/default/snipcart.js"></script>
    <div hidden id="snipcart" data-api-key="YOUR_PUBLIC_API_KEY"></div>
  </body>
</html>
```


https://docs.snipcart.com/v3/setup/products
add to cart button
```
<button class="snipcart-add-item"
  data-item-id="starry-night"
  data-item-price="79.99"
  data-item-url="/paintings/starry-night"
  data-item-description="High-quality replica of The Starry Night by the Dutch post-impressionist painter Vincent van Gogh."
  data-item-image="/assets/images/starry-night.jpg"
  data-item-name="The Starry Night">
  Add to cart
</button>
```

## BTCPayServer integration

live demo:
https://mainnet.demo.btcpayserver.org/

scripts and stylesheet


buy button with default endpoint:
```
<style type="text/css"> .btcpay-form { display: inline-flex; align-items: center; justify-content: center; } .btcpay-form--inline { flex-direction: row; } .btcpay-form--block { flex-direction: column; } .btcpay-form--inline .submit { margin-left: 15px; } .btcpay-form--block select { margin-bottom: 10px; } .btcpay-form .btcpay-custom-container{ text-align: center; }.btcpay-custom { display: flex; align-items: center; justify-content: center; } .btcpay-form .plus-minus { cursor:pointer; font-size:25px; line-height: 25px; background: #DFE0E1; height: 30px; width: 45px; border:none; border-radius: 60px; margin: auto 5px; display: inline-flex; justify-content: center; } .btcpay-form select { -moz-appearance: none; -webkit-appearance: none; appearance: none; color: currentColor; background: transparent; border:1px solid transparent; display: block; padding: 1px; margin-left: auto; margin-right: auto; font-size: 11px; cursor: pointer; } .btcpay-form select:hover { border-color: #ccc; } #btcpay-input-price { -moz-appearance: none; -webkit-appearance: none; border: none; box-shadow: none; text-align: center; font-size: 25px; margin: auto; border-radius: 5px; line-height: 35px; background: #fff; } </style>
<form method="POST"  action="http://127.0.0.1:23000/api/v1/invoices" class="btcpay-form btcpay-form--block">
  <input type="hidden" name="storeId" value="5VrWYtRDGfraCA9A9xfsNizEt6moU5hZyG1pShrJS3Cv" />
  <input type="hidden" name="price" value="10" />
  <input type="hidden" name="currency" value="USD" />
  <input type="image" class="submit" name="submit" src="http://127.0.0.1:23000/img/paybutton/pay.svg" style="width:209px" alt="Pay with BtcPay, Self-Hosted Bitcoin Payment Processor">
</form>
```

buy button with app endpoint and product
```
<style type="text/css"> .btcpay-form { display: inline-flex; align-items: center; justify-content: center; } .btcpay-form--inline { flex-direction: row; } .btcpay-form--block { flex-direction: column; } .btcpay-form--inline .submit { margin-left: 15px; } .btcpay-form--block select { margin-bottom: 10px; } .btcpay-form .btcpay-custom-container{ text-align: center; }.btcpay-custom { display: flex; align-items: center; justify-content: center; } .btcpay-form .plus-minus { cursor:pointer; font-size:25px; line-height: 25px; background: #DFE0E1; height: 30px; width: 45px; border:none; border-radius: 60px; margin: auto 5px; display: inline-flex; justify-content: center; } .btcpay-form select { -moz-appearance: none; -webkit-appearance: none; appearance: none; color: currentColor; background: transparent; border:1px solid transparent; display: block; padding: 1px; margin-left: auto; margin-right: auto; font-size: 11px; cursor: pointer; } .btcpay-form select:hover { border-color: #ccc; } #btcpay-input-price { -moz-appearance: none; -webkit-appearance: none; border: none; box-shadow: none; text-align: center; font-size: 25px; margin: auto; border-radius: 5px; line-height: 35px; background: #fff; } </style>
<form method="POST"  action="http://127.0.0.1:23000/apps/42U8i51gFZF7reRf71hZ26wa7Qx3/pos" class="btcpay-form btcpay-form--block">
  <input type="hidden" name="storeId" value="5VrWYtRDGfraCA9A9xfsNizEt6moU5hZyG1pShrJS3Cv" />
  <input type="hidden" name="choiceKey" value="cap1500uf6p3e" />
  <input type="hidden" name="amount" value="10" />
  <input type="image" class="submit" name="submit" src="http://127.0.0.1:23000/img/paybutton/pay.svg" style="width:209px" alt="Pay with BtcPay, Self-Hosted Bitcoin Payment Processor">
</form>
```
