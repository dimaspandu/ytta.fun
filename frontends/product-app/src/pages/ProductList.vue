<template>
  <div class="product-list">
    <h1 class="product-list__title">Product List</h1>

    <div v-if="filteredProducts.length === 0" class="product-list__empty">
      No products available.
    </div>

    <div class="product-list__cards">
      <ProductCard
        v-for="product in filteredProducts"
        :key="product.id"
        :item="product"
      />
    </div>
  </div>
</template>

<script>
import { reactive, computed, onMounted } from 'vue';
import ProductCard from '../components/ProductCard.vue';
import { sessionStore } from '../store/session';

export default {
  name: 'ProductList',
  components: { ProductCard },
  setup() {
    const products = reactive([]);

    onMounted(() => {
      products.push(
        {
          id: 1,
          name: 'Product A',
          price: 29.99,
          seller: 'Seller One',
          expiresAt: new Date(Date.now() + 3600 * 1000).toISOString(),
          image: 'https://picsum.photos/300/200?random=1'
        },
        {
          id: 2,
          name: 'Product B',
          price: 49.5,
          seller: 'Seller Two',
          expiresAt: new Date(Date.now() + 7200 * 1000).toISOString(),
          image: 'https://picsum.photos/300/200?random=2'
        },
        {
          id: 3,
          name: 'Product C',
          price: 15.0,
          seller: 'Seller Three',
          expiresAt: new Date(Date.now() + 10800 * 1000).toISOString(),
          image: 'https://picsum.photos/300/200?random=3'
        }
      );

      sessionStore.removeExpired();
    });

    const filteredProducts = computed(() => {
      const now = new Date();
      return products.filter(product => new Date(product.expiresAt) > now);
    });

    return { filteredProducts };
  }
};
</script>

<style lang="scss" scoped>
@use '../styles/main.scss' as *;

.product-list {
  padding: 1rem;
  margin: 0 auto;
  max-width: 1200px;
}

.product-list__title {
  color: $primary-color;
  margin-bottom: 1rem;
  font-size: 2rem;
  font-weight: 700;
}

.product-list__cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 1.2rem;
}

.product-list__empty {
  color: #888;
  font-style: italic;
  margin-top: 2rem;
  text-align: center;
  font-size: 1.3rem;
}

@media (max-width: 600px) {
  .product-list {
    padding: 1.25rem;
  }
  .product-list__cards {
    grid-template-columns: 1fr;
  }
}
</style>
