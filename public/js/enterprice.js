// product model
window.Product = Backbone.Model.extend({});

// products collection
window.Products = Backbone.Collection.extend({
  model: Product
});

// product view
window.ProductView = Backbone.View.extend({
  template: _.template("<%= name %>, <%= price %>"),
  tagName: 'p',
  render: function() {
    this.$el.html(this.template(this.model.toJSON()));
    return this;
  }
});

// initialize router
window.Enterprice = new(Backbone.Router.extend({
  routes: {
    "": "index",
    "products/:id": "show"
  },

  index: function () {
    console.log("Hello from Backbone");

    var product = new Product({name: "Bread", price: "3.40"});
    var productView = new ProductView({model: product});
    productView.render();
    $('#products').html(productView.el);
  },

  show: function() {
  },

  start: function(){
    Backbone.history.start();
  }
}));

// start Backbone app
$(function(){
  Enterprice.start()
});