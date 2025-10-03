<template>
  <div class="product-detail">
    <router-link class="product-detail__back" to="/">
      <span class="product-detail__back-icon">
        <svg width="800px" height="800px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M5 12H19M5 12L11 6M5 12L11 18" stroke="#000000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
        </svg>
      </span>
      Back to list
    </router-link>

    <div class="product-detail__container">
      <div class="product-detail__carousel">
        <div class="carousel-track" :style="{ transform: `translateX(-${currentIndex * 100}%)` }">
          <img
            v-for="(img, index) in product.images"
            :key="index"
            :src="img"
            @error="onImageError($event, index)"
            :alt="product.name + ' ' + (index + 1)"
          />
        </div>
        <button class="carousel-btn prev" @click="prevImage">&#10094;</button>
        <button class="carousel-btn next" @click="nextImage">&#10095;</button>
      </div>

      <div class="product-detail__info">
        <h1 class="product-detail__title">{{ product.name }}</h1>

        <div class="product-detail__field">
          <strong>Seller:</strong>
          <span>{{ product.seller }}</span>
        </div>
        <div class="product-detail__field">
          <strong>Price:</strong>
          <span>{{ product.price ? '$' + product.price.toFixed(2) : 'N/A' }}</span>
        </div>
        <div class="product-detail__field">
          <strong>Expires at:</strong>
          <span>{{ new Date(product.expiresAt).toLocaleString() }}</span>
        </div>

        <div class="product-detail__order">
          <label>
            Quantity:
            <input type="number" v-model.number="quantity" min="1" />
          </label>

          <label>
            Variant:
            <select v-model="variant">
              <option v-for="v in product.variants" :key="v" :value="v">{{ v }}</option>
            </select>
          </label>

          <label>
            Notes:
            <textarea v-model="notes" placeholder="Add notes for your order"></textarea>
          </label>

          <button class="product-detail__order-btn">Order</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { reactive, ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';

export default {
  name: 'ProductDetail',
  setup() {
    const route = useRoute();
    const product = reactive({
      id: null,
      name: '',
      price: 0,
      seller: '',
      expiresAt: '',
      images: [],
      variants: []
    });

    const quantity = ref(1);
    const variant = ref('');
    const notes = ref('');
    const currentIndex = ref(0);
    const fallbackImage = 'https://picsum.photos/400/300?blur=5';

    const onImageError = (event) => {
      event.target.src = fallbackImage;
    };

    const nextImage = () => {
      currentIndex.value = (currentIndex.value + 1) % product.images.length;
    };

    const prevImage = () => {
      currentIndex.value =
        (currentIndex.value - 1 + product.images.length) % product.images.length;
    };

    onMounted(() => {
      const id = Number(route.params.id);
      product.id = id;
      product.name = `Product ${String.fromCharCode(64 + id)}`;
      product.price = 20 * id + 9.99;
      product.seller = `Seller ${id}`;
      product.expiresAt = new Date(Date.now() + 3600 * 1000 * id).toISOString();
      product.images = [
        `https://picsum.photos/400/300?random=${id * 1}`,
        `https://picsum.photos/400/300?random=${id * 2}`,
        `https://picsum.photos/400/300?random=${id * 3}`
      ];
      product.variants = ['Red', 'Blue', 'Green'];
      variant.value = product.variants[0];
    });

    return {
      product,
      quantity,
      variant,
      notes,
      currentIndex,
      onImageError,
      nextImage,
      prevImage
    };
  }
};
</script>

<style lang="scss" scoped>
@use '../styles/main.scss' as *;

.product-detail {
  padding: 2.5rem;
  max-width: 900px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 1rem;

  &__back {
    display: inline-flex;
    align-items: center;
    gap: 0.4rem;
    color: $primary-color;
    text-decoration: none;
    font-size: 1.2rem;
    font-weight: 600;

    &-icon {
      display: flex;
      align-items: center;

      svg {
        width: 1rem;
        height: 1rem;
      }
    }

    &:hover {
      text-decoration: underline;
    }
  }

  &__container {
    display: flex;
    gap: 2rem;
    flex-wrap: wrap;
  }

  &__carousel {
    position: relative;
    flex: 1 1 250px;
    overflow: hidden;
    border-radius: 12px;
    height: 250px;

    .carousel-track {
      display: flex;
      transition: transform 0.4s ease;
      img {
        width: 100%;
        flex-shrink: 0;
        border-radius: 12px;
        object-fit: cover;
      }
    }

    .carousel-btn {
      position: absolute;
      top: 50%;
      transform: translateY(-50%);
      background: rgba(0,0,0,0.4);
      color: #fff;
      border: none;
      padding: 0.5rem 1rem;
      font-size: 1.5rem;
      cursor: pointer;
      border-radius: 6px;

      &.prev { left: 10px; }
      &.next { right: 10px; }

      &:hover { background: rgba(0,0,0,0.6); }
    }
  }

  &__info {
    flex: 1 1 300px;
    display: flex;
    flex-direction: column;
    gap: 0.8rem;
  }

  &__title {
    color: $primary-color;
    font-size: 2rem;
    font-weight: 700;
    margin: 0.5rem 0;
  }

  &__field {
    display: flex;
    flex-direction: column;
    font-size: 1.3rem;
    color: #333;
    gap: 0.2rem;
    strong {
      font-weight: 700;
    }
  }

  &__order {
    display: flex;
    flex-direction: column;
    gap: 0.8rem;
    background: rgba(255,255,255,0.15);
    padding: 1rem;
    border-radius: 16px;
    backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0,0,0,0.12);

    label {
      font-size: 1.2rem;
      display: flex;
      flex-direction: column;
      gap: 0.3rem;

      input, select, textarea {
        padding: 0.5rem;
        font-size: 1.2rem;
        border-radius: 8px;
        border: 1px solid #ccc;
      }

      textarea { min-height: 60px; resize: vertical; }
    }

    &__order-btn, button.product-detail__order-btn {
      margin-top: 0.5rem;
      background-color: $primary-color;
      color: #fff;
      border: none;
      border-radius: 10px;
      font-size: 1.2rem;
      font-weight: 600;
      padding: 0.7rem 1.2rem;
      cursor: pointer;
      transition: background-color 0.3s ease, transform 0.2s ease;

      &:hover {
        background-color: $primary-hover-color;
        transform: scale(1.05);
      }
    }
  }

  @media (max-width: 480px) {
    padding: 1.25rem;
  }
}

@media (max-width: 768px) {
  .product-detail {
    &__container {
      flex-direction: column;
    }

    &__carousel {
      height: 220px;

      .carousel-track {
        flex-direction: row;
      }
    }
  }
}
</style>
