export type RegisterFormData = {
  role: "plug" | "user";
  username: string;
  name: string;
  email: string;
  password: string;
  confirmPassword: string;
};

export type RequestFormData = {
  requestType: "shipping" | "delivery";
  title: string;
  description: string;
  place: string;
  phoneNumber: string;
};
