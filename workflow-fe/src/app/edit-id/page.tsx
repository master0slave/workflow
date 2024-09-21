import { useRouter } from "next/navigation";

const EditItem = () => {
    const router = useRouter();
    const { id } = router.query;
    const [item, setItem] = useState(null)

    useEffect(() => {
        const fetchItem = async () => {
            const response = await fetch(`/api/items/${id}`)
            const data = await response.json()
            setItem(data)
        }
        fetchItem()
    }, [id]) 
    return (
}