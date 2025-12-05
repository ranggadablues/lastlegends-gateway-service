window.onload = function () {
  //<editor-fold desc="Changeable Configuration Block">

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  const ui = SwaggerUIBundle({
    urls: [
      { url: "/swagger/product.swagger.json", name: "Product API" },
      { url: "/swagger/user.swagger.json", name: "User API" }
    ],
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    layout: "StandaloneLayout"
  })
  window.ui = ui

  //</editor-fold>
};
