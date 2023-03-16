import {makeAutoObservable} from "mobx";

interface Modal {
    open: boolean
    body: JSX.Element | null
}

export default class ModalStore {
    modal: Modal = {
        open: false,
        body: null,
    }

    constructor() {
        makeAutoObservable(this)
    }

    openModal = (content: JSX.Element) => {
        this.modal.body = content
        this.modal.open = true
    }

    closeModal = () => {
        this.modal.open = false
        this.modal.body = null
    }
}