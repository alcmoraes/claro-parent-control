package messages

func English() map[string]string {
    return map[string]string{
        "BTN_REAUTH_ROUTER": "Reauth router",
        "BTN_ALLOW_MAC": "Allow MAC",
        "BTN_BLOCK_MAC": "Block MAC",
        "BTN_CLEAR_BLACKLIST": "Clear blacklist",
        "BTN_RESTART": "Reiniciar RPi",
        
        "MSG_WELCOME": "Hello!\nUse: \n\n/login PASSWORD\n\nto authenticate before using it",
        "MSG_UNAUTHENTICATED": "You are not authenticated. Use /login PASSWORD to authenticate",
        "MSG_AUTHENTICATED": "You are authenticated",
        "MSG_WRONG_PASSWORD": "Wrong password",
        "MSG_REAUTH_FAILED": "Failed to reauth router",
        "MSG_REAUTH_SUCCESS": "Router reauthenticated successfully",
        "MSG_FAILED_TO_GET_BLACKLIST": "Failed to get blacklist",
        "MSG_CHOOSE_DEVICE": "Choose a device",
        "MSG_NO_DEVICES_BLOCKED": "No devices blocked",
        "MSG_FAILED_TO_CLEAR_BLACKLIST": "Failed to clear blacklist",
        "MSG_BLACKLIST_CLEARED": "Blacklist cleared",
        "MSG_FAILED_TO_UNBLOCK_DEVICE": "Failed to unblock device",
        "MSG_DEVICE_UNBLOCKED": "Device unblocked",
        "MSG_FAILED_TO_BLOCK_DEVICE": "Failed to block device",
        "MSG_DEVICE_BLOCKED": "Device blocked",
    }
}