export interface User {
  id: number;
  email: string;
  created_at: string;
}

export interface LoginResponse {
  token: string;
}

export interface Application {
  id: number;
  company: string;
  position: string;
  job_url: string | null;
  applied_at: string | null;
  notes: string | null;
  current_status: string;
  created_at: string;
  updated_at: string;
}

export interface CreateApplicationRequest {
  company: string;
  position: string;
  job_url?: string;
  applied_at?: string;
  notes?: string;
}

export interface UpdateStatusRequest {
  status: string;
}