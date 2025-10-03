import { reactive } from 'vue';

export const sessionStore = reactive({
  sessions: [],
  addSession(session) {
    this.sessions.push(session)
  },
  removeExpired() {
    const now = new Date()
    this.sessions = this.sessions.filter(s => new Date(s.expiresAt) > now)
  }
});
