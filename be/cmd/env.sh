# env.sh

# Port for the application
export APP_PORT=4100

# URLs for payment notifications and redirects
export WEBHOOK_NOTIFICATION_URL="https://your-notification-url.com"
export RETURN_SUCCESS_URL="https://your-return-success-url.com"
export RETURN_FAILURE_URL="https://your-return-failure-url.com"
export RETURN_CANCEL_URL="https://your-return-cancel-url.com"
export RETURN_PENDING_URL="https://your-return-pending-url.com"

# Terminal ID for emerchantpay
export EMERCHANTPAY_TERMINAL_ID="your_terminal_id_here"

# Emerchantpay credentials
export EMERCHANTPAY_USERNAME="your_username_here"
export EMERCHANTPAY_PASSWORD="your_password_here"

# Emerchantpay API URLS
export EMERCHANTPAY_WPF_API_URL="https://staging.wpf.emerchantpay.net/en/wpf"
