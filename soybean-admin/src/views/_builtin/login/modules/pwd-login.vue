<script setup lang="ts">
import { computed, reactive } from 'vue';
import { $t } from '@/locales';
import { useNaiveForm } from '@/hooks/common/form';
import { useAuthStore } from '@/store/modules/auth';

defineOptions({
  name: 'PwdLogin'
});

const authStore = useAuthStore();
const { formRef, validate } = useNaiveForm();

const model: Api.Form.LoginForm = reactive({
  username: '',
  password: '',
  code: '',
  captchaId: ''
});

const rules = computed<Record<keyof Api.Form.LoginForm, App.Global.FormRule[]>>(() => {
  return {
    username: [
      {
        required: true,
        message: 'Kullanıcı adı boş bırakılamaz'
      }
    ],
    password: [
      {
        required: true,
        message: 'Parola boş bırakılamaz'
      }
    ],
    code: [],
    captchaId: []
  };
});

async function handleSubmit() {
  await validate();
  await authStore.login(model);
}
</script>

<template>
  <NForm ref="formRef" :model="model" :rules="rules" size="large" :show-label="false">
    <NFormItem path="username">
      <NInput v-model:value="model.username" :placeholder="$t('page.login.common.userNamePlaceholder')" />
    </NFormItem>
    <NFormItem path="password">
      <NInput
        v-model:value="model.password"
        type="password"
        show-password-on="click"
        :placeholder="$t('page.login.common.passwordPlaceholder')"
      />
    </NFormItem>
    <NSpace vertical :size="24">
      <div class="flex-y-center justify-between">
        <NCheckbox>{{ $t('page.login.pwdLogin.rememberMe') }}</NCheckbox>
      </div>
      <NButton
        attr-type="submit"
        type="primary"
        size="large"
        round
        block
        :loading="authStore.loginLoading"
        @click="handleSubmit"
      >
        {{ $t('common.confirm') }}
      </NButton>
    </NSpace>
  </NForm>
</template>

<style scoped></style>
