import type {
  LoginResponse,
  User,
  Application,
  CreateApplicationRequest,
  UpdateStatusRequest,
} from "@/types";


const API_URL = process.env.NEXT_PUBLIC_API_URL;

class ApiError extends Error {
  constructor(public status: number, message: string) {
    super(message);
  }
}

async function apiFetch<T>(
  path: string,
  options: RequestInit = {}
): Promise<T> {
  const token = typeof window !== "undefined" ? localStorage.getItem("token") : null;

  const headers: HeadersInit = {
    "Content-Type": "application/json",
    ...(token ? { Authorization: `Bearer ${token}` } : {}),
    ...options.headers,
  };

  const res = await fetch(`${API_URL}${path}`, { ...options, headers });

  if (!res.ok) {
    const body = await res.json().catch(() => ({ error: "Terjadi kesalahan" }));
    throw new ApiError(res.status, body.error || "Terjadi kesalahan");
  }

  if (res.status === 204) {
    return undefined as T;
  }

  return res.json();
}

export function login(email: string, password: string) {
  return apiFetch<LoginResponse>("/auth/login", {
    method: "POST",
    body: JSON.stringify({ email, password }),
  });
}

export function register(email: string, password: string) {
  return apiFetch<User>("/auth/register", {
    method: "POST",
    body: JSON.stringify({ email, password }),
  });
}

export function forgotPassword(email: string) {
  return apiFetch<{ message: string }>("/auth/forgot-password", {
    method: "POST",
    body: JSON.stringify({ email }),
  });
}

export function resetPassword(token: string, password: string) {
  return apiFetch<{ message: string }>("/auth/reset-password", {
    method: "POST",
    body: JSON.stringify({ token, password }),
  });
}

export function getApplications() {
  return apiFetch<Application[]>("/applications");
}

export function getApplication(id: number) {
  return apiFetch<Application>(`/applications/${id}`);
}

export function createApplication(data: CreateApplicationRequest) {
  return apiFetch<Application>("/applications", {
    method: "POST",
    body: JSON.stringify(data),
  });
}

export function updateApplicationStatus(id: number, data: UpdateStatusRequest) {
  return apiFetch<Application>(`/applications/${id}/status`, {
    method: "PATCH",
    body: JSON.stringify(data),
  });
}

export function deleteApplication(id: number) {
  return apiFetch<void>(`/applications/${id}`, {
    method: "DELETE",
  });
}