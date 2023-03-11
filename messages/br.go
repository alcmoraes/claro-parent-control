package messages

func Brazilian() map[string]string {
    return map[string]string{
        "BTN_REAUTH_ROUTER": "Relogar Modem",
        "BTN_ALLOW_MAC": "Permitir MAC",
        "BTN_BLOCK_MAC": "Bloquear MAC",
        "BTN_CLEAR_BLACKLIST": "Limpar bloqueios",
        "BTN_RESTART": "Reiniciar RPi",

        "MSG_WELCOME": "Olá!\nUse: \n\n/login PASSWORD\n\npara se autenticar antes de usar o bot",
        "MSG_UNAUTHENTICATED": "Você não está autenticado. Use /login PASSWORD para se autenticar",
        "MSG_AUTHENTICATED": "Você está autenticado",
        "MSG_WRONG_PASSWORD": "Senha incorreta",
        "MSG_REAUTH_FAILED": "Falha ao relogar modem",
        "MSG_REAUTH_SUCCESS": "Modem relogado com sucesso",
        "MSG_FAILED_TO_GET_BLACKLIST": "Falha ao obter lista de bloqueios",
        "MSG_CHOOSE_DEVICE": "Escolha um dispositivo",
        "MSG_NO_DEVICES_BLOCKED": "Nenhum dispositivo bloqueado",
        "MSG_FAILED_TO_CLEAR_BLACKLIST": "Falha ao limpar lista de bloqueios",
        "MSG_BLACKLIST_CLEARED": "Lista de bloqueios limpa",
        "MSG_FAILED_TO_UNBLOCK_DEVICE": "Falha ao desbloquear dispositivo",
        "MSG_DEVICE_UNBLOCKED": "Dispositivo desbloqueado",
        "MSG_FAILED_TO_BLOCK_DEVICE": "Falha ao bloquear dispositivo",
        "MSG_DEVICE_BLOCKED": "Dispositivo bloqueado",
        "MSG_UNKNOWN_COMMAND": "Comando desconhecido",
    }
}