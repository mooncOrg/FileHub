import request from '@/utils/request'

/**
 * 上传文件
 * @param formData 包含 file 的表单数据
 * @param onProgress 进度回调函数
 */
export const uploadFileApi = (
    formData: FormData,
    onProgress: (percent: number) => void
) => {
    return request.post('/upload', formData, {
        // 处理 Axios 上传进度
        onUploadProgress: (progressEvent) => {
            if (progressEvent.total) {
                const percent = Math.round((progressEvent.loaded * 100) / progressEvent.total)
                onProgress(percent)
            }
        }
    })
}