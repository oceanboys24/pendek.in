import { useMutation, useQueryClient } from "@tanstack/react-query";
import { AxiosError } from "axios";
import axios from "axios";
import { toast } from "sonner";

export const useShortLink = () => {
  const queryClient = useQueryClient();
  const { mutateAsync, isPending, isSuccess, data } = useMutation<
    { message: string; short_link: string },
    AxiosError,
    { url: string }
  >({
    mutationKey: ["short-url"],
    mutationFn: async (data: { url: string }) => {
      const response = await axios.post("https://api.alfiandirizki.com/", data);
      return response.data;
    },
    onSuccess: async (res) => {
      await queryClient.setQueryData(["short-url"], res.short_link);
      toast.success(res.message);
    },
    onError: (error) => {
      if (error) {
        toast.error("err");
      }
    },
  });
  const getData = queryClient.getQueryData(["short-url"]);

  return { mutateAsync, isPending, getData, isSuccess, data };
};
