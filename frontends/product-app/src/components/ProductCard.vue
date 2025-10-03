<template>
  <div class="product-card">
    <div class="product-card__image-wrapper">
      <img
        class="product-card__image"
        :src="item.image || fallbackImage"
        @error="onImageError"
        :alt="item.name"
      />
    </div>
    <div class="product-card__info">
      <h3 class="product-card__name">
        <router-link :to="{ name: 'ProductDetail', params: { id: item.id } }">
          {{ item.name }}
        </router-link>
      </h3>
      <p class="product-card__seller">
        Seller: {{ item.seller || 'Unknown' }}
      </p>
      <p class="product-card__price">
        Price: {{ item.price ? '$' + item.price.toFixed(2) : 'N/A' }}
      </p>
      <p class="product-card__expiry">
        Expires: {{ new Date(item.expiresAt).toLocaleString() }}
      </p>
      <div class="product-card__buttons">
        <button class="product-card__order-btn">Order</button>
        <router-link
          :to="{ name: 'ProductDetail', params: { id: item.id } }"
          class="product-card__detail-btn"
        >
          Lihat Detail
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ProductCard',
  props: {
    item: Object
  },
  data() {
    return {
      fallbackImage: 'https://picsum.photos/300/200?blur=5'
    };
  },
  methods: {
    onImageError(event) {
      event.target.src = this.fallbackImage;
    }
  }
};
</script>

<style lang="scss" scoped>
@use '../styles/main.scss' as *;

.product-card {
  display: flex;
  flex-direction: column;
  border-radius: 16px;
  padding: 0;
  overflow: hidden;
  background: rgba(255, 255, 255, 0.15);
  backdrop-filter: blur(10px);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.12);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
  cursor: pointer;

  &:hover {
    transform: translateY(-5px);
    box-shadow: 0 16px 48px rgba(0, 0, 0, 0.2);
  }

  &__image-wrapper {
    width: 100%;
    height: 220px;
    overflow: hidden;
    border-bottom: 1px solid rgba(0, 0, 0, 0.1);
  }

  &__image {
    width: 100%;
    height: 100%;
    object-fit: cover;
    display: block;
  }

  &__info {
    padding: 1.2rem;
    display: flex;
    flex-direction: column;
    gap: 0.6rem;
  }

  &__name {
    font-weight: 700;
    color: #222;
    margin: 0;

    a {
      font-size: 1.5rem;
      color: inherit;
      text-decoration: none;

      &:hover {
        color: $primary-color;
      }
    }
  }

  &__seller {
    font-size: 1.2rem;
    color: #444;
    margin: 0;
  }

  &__price {
    font-size: 1.3rem;
    font-weight: 600;
    color: #222;
    margin: 0;
  }

  &__expiry {
    color: #555;
    font-size: 1.2rem;
    margin: 0;
  }

  &__buttons {
    display: flex;
    gap: 0.8rem;
    margin-top: 0.8rem;
  }

  &__order-btn,
  &__detail-btn {
    padding: 0.7rem 1.2rem;
    font-size: 1.2rem;
    font-weight: 600;
    border-radius: 10px;
    cursor: pointer;
    transition: background-color 0.3s ease, transform 0.2s ease;
    text-decoration: none;
    text-align: center;
  }

  &__order-btn {
    background-color: $primary-color;
    color: #fff;
    border: none;

    &:hover {
      background-color: $primary-hover-color;
      transform: scale(1.05);
    }
  }

  &__detail-btn {
    background-color: transparent;
    color: $primary-color;
    border: 2px solid $primary-color;

    &:hover {
      background-color: $primary-color;
      color: #fff;
      transform: scale(1.05);
    }
  }
}
</style>
